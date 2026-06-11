#include <errno.h>
#include <stdint.h>
#include <string.h>

#include <zephyr/drivers/i2c.h>
#include <zephyr/kernel.h>
#include <zephyr/ztest.h>

#define BADGESNAKE_I2C_NODE DT_ALIAS(badgesnake_i2c)

#if !DT_NODE_EXISTS(BADGESNAKE_I2C_NODE)
#error "badgesnake-i2c alias is required"
#endif

#define BADGESNAKE_I2C_ADDR 0x42
#define BADGESNAKE_PROTOCOL_VERSION 0x01
#define BADGESNAKE_HEADER_SIZE 6
#define BADGESNAKE_MAX_FRAME_SIZE 1024

enum badgesnake_token {
	BADGESNAKE_TOKEN_INFO = 0x01,
	BADGESNAKE_TOKEN_START = 0x02,
	BADGESNAKE_TOKEN_MOVE = 0x03,
	BADGESNAKE_TOKEN_END = 0x04,
};

enum badgesnake_status {
	BADGESNAKE_STATUS_SUCCESS = 0x00,
	BADGESNAKE_STATUS_BAD_REQUEST = 0x01,
	BADGESNAKE_STATUS_UNSUPPORTED = 0x02,
	BADGESNAKE_STATUS_INTERNAL = 0x03,
};

struct badgesnake_target_state {
	struct i2c_target_config config;
	uint8_t rx_buf[BADGESNAKE_MAX_FRAME_SIZE];
	size_t rx_len;
	uint8_t tx_buf[BADGESNAKE_MAX_FRAME_SIZE];
	size_t tx_len;
	size_t tx_pos;
	bool stop_seen;
};

static const struct device *const badgesnake_i2c = DEVICE_DT_GET(BADGESNAKE_I2C_NODE);

static const char info_response_json[] =
	"{"
	"\"apiversion\":\"1\","
	"\"author\":\"BadgeSnake\","
	"\"color\":\"#2f8f6b\","
	"\"head\":\"pixel\","
	"\"tail\":\"bolt\","
	"\"version\":\"0.1.0\""
	"}";

static const char start_request_json[] =
	"{"
	"\"game\":{\"id\":\"badgesnake-demo-game\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\","
	"\"settings\":{\"foodSpawnChance\":15,\"minimumFood\":1,\"hazardDamagePerTurn\":14,"
	"\"royale\":{\"shrinkEveryNTurns\":0},"
	"\"squad\":{\"allowBodyCollisions\":false,\"sharedElimination\":false,"
	"\"sharedHealth\":false,\"sharedLength\":false}}},\"timeout\":500,\"source\":\"badgesnake-sim\"},"
	"\"turn\":0,"
	"\"board\":{\"height\":11,\"width\":11,\"food\":[{\"x\":5,\"y\":5}],\"hazards\":[],"
	"\"snakes\":[{\"id\":\"player-1\",\"name\":\"Player 1\",\"health\":100,"
	"\"body\":[{\"x\":1,\"y\":1},{\"x\":1,\"y\":0},{\"x\":1,\"y\":0}],"
	"\"head\":{\"x\":1,\"y\":1},\"length\":3,\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#2f8f6b\",\"head\":\"pixel\",\"tail\":\"bolt\"}}]},"
	"\"you\":{\"id\":\"player-1\",\"name\":\"Player 1\",\"health\":100,"
	"\"body\":[{\"x\":1,\"y\":1},{\"x\":1,\"y\":0},{\"x\":1,\"y\":0}],"
	"\"head\":{\"x\":1,\"y\":1},\"length\":3,\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#2f8f6b\",\"head\":\"pixel\",\"tail\":\"bolt\"}}"
	"}";

static const char move_request_json[] =
	"{"
	"\"game\":{\"id\":\"badgesnake-demo-game\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\","
	"\"settings\":{\"foodSpawnChance\":15,\"minimumFood\":1,\"hazardDamagePerTurn\":14,"
	"\"royale\":{\"shrinkEveryNTurns\":0},"
	"\"squad\":{\"allowBodyCollisions\":false,\"sharedElimination\":false,"
	"\"sharedHealth\":false,\"sharedLength\":false}}},\"timeout\":500,\"source\":\"badgesnake-sim\"},"
	"\"turn\":3,"
	"\"board\":{\"height\":11,\"width\":11,\"food\":[{\"x\":5,\"y\":5},{\"x\":8,\"y\":2}],"
	"\"hazards\":[],"
	"\"snakes\":[{\"id\":\"player-1\",\"name\":\"Player 1\",\"health\":97,"
	"\"body\":[{\"x\":3,\"y\":1},{\"x\":2,\"y\":1},{\"x\":1,\"y\":1}],"
	"\"head\":{\"x\":3,\"y\":1},\"length\":3,\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#2f8f6b\",\"head\":\"pixel\",\"tail\":\"bolt\"}},"
	"{\"id\":\"player-2\",\"name\":\"Player 2\",\"health\":97,"
	"\"body\":[{\"x\":7,\"y\":9},{\"x\":7,\"y\":10},{\"x\":7,\"y\":10}],"
	"\"head\":{\"x\":7,\"y\":9},\"length\":3,\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#d96c3f\",\"head\":\"pixel\",\"tail\":\"bolt\"}}]},"
	"\"you\":{\"id\":\"player-1\",\"name\":\"Player 1\",\"health\":97,"
	"\"body\":[{\"x\":3,\"y\":1},{\"x\":2,\"y\":1},{\"x\":1,\"y\":1}],"
	"\"head\":{\"x\":3,\"y\":1},\"length\":3,\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#2f8f6b\",\"head\":\"pixel\",\"tail\":\"bolt\"}}"
	"}";

static const char move_response_json[] =
	"{"
	"\"move\":\"up\","
	"\"shout\":\"badgesnake\""
	"}";

static const char end_request_json[] =
	"{"
	"\"game\":{\"id\":\"badgesnake-demo-game\",\"ruleset\":{\"name\":\"standard\",\"version\":\"cli\","
	"\"settings\":{\"foodSpawnChance\":15,\"minimumFood\":1,\"hazardDamagePerTurn\":14,"
	"\"royale\":{\"shrinkEveryNTurns\":0},"
	"\"squad\":{\"allowBodyCollisions\":false,\"sharedElimination\":false,"
	"\"sharedHealth\":false,\"sharedLength\":false}}},\"timeout\":500,\"source\":\"badgesnake-sim\"},"
	"\"turn\":12,"
	"\"board\":{\"height\":11,\"width\":11,\"food\":[{\"x\":5,\"y\":5}],\"hazards\":[],"
	"\"snakes\":[{\"id\":\"player-1\",\"name\":\"Player 1\",\"health\":0,"
	"\"body\":[{\"x\":4,\"y\":4}],\"head\":{\"x\":4,\"y\":4},\"length\":1,"
	"\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#2f8f6b\",\"head\":\"pixel\",\"tail\":\"bolt\"}}]},"
	"\"you\":{\"id\":\"player-1\",\"name\":\"Player 1\",\"health\":0,"
	"\"body\":[{\"x\":4,\"y\":4}],\"head\":{\"x\":4,\"y\":4},\"length\":1,"
	"\"latency\":\"0\",\"shout\":\"\",\"squad\":\"\","
	"\"customizations\":{\"color\":\"#2f8f6b\",\"head\":\"pixel\",\"tail\":\"bolt\"}}"
	"}";

static struct badgesnake_target_state target_state;
static bool target_registered;

static uint16_t badgesnake_u16le(const uint8_t *buf)
{
	return (uint16_t)buf[0] | ((uint16_t)buf[1] << 8);
}

static int badgesnake_encode_frame(uint8_t code, const uint8_t *payload, size_t payload_len,
				   uint8_t *buf, size_t *buf_len)
{
	if (payload_len > UINT16_MAX) {
		return -EMSGSIZE;
	}
	if ((BADGESNAKE_HEADER_SIZE + payload_len) > BADGESNAKE_MAX_FRAME_SIZE) {
		return -EMSGSIZE;
	}

	buf[0] = BADGESNAKE_PROTOCOL_VERSION;
	buf[1] = code;
	buf[2] = 0;
	buf[3] = 0;
	buf[4] = payload_len & 0xff;
	buf[5] = (payload_len >> 8) & 0xff;
	if (payload_len > 0) {
		memcpy(&buf[BADGESNAKE_HEADER_SIZE], payload, payload_len);
	}
	*buf_len = BADGESNAKE_HEADER_SIZE + payload_len;

	return 0;
}

static int badgesnake_prepare_response(struct badgesnake_target_state *state, uint8_t status,
				       const uint8_t *payload, size_t payload_len)
{
	state->tx_pos = 0;
	return badgesnake_encode_frame(status, payload, payload_len, state->tx_buf, &state->tx_len);
}

static int badgesnake_prepare_json_response(struct badgesnake_target_state *state, uint8_t status,
					    const char *json)
{
	return badgesnake_prepare_response(state, status, (const uint8_t *)json, strlen(json));
}

static bool badgesnake_payload_matches(const uint8_t *payload, size_t payload_len, const char *expected)
{
	size_t expected_len = strlen(expected);

	if (payload_len != expected_len) {
		return false;
	}

	return memcmp(payload, expected, payload_len) == 0;
}

static int badgesnake_dispatch_request(struct badgesnake_target_state *state)
{
	const uint8_t *payload;
	size_t payload_len;

	if (state->rx_len < BADGESNAKE_HEADER_SIZE) {
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
							"{\"error\":\"frame too short\"}");
	}

	if (state->rx_buf[0] != BADGESNAKE_PROTOCOL_VERSION) {
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
							"{\"error\":\"unsupported version\"}");
	}

	payload_len = badgesnake_u16le(&state->rx_buf[4]);
	if ((BADGESNAKE_HEADER_SIZE + payload_len) != state->rx_len) {
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
							"{\"error\":\"length mismatch\"}");
	}

	payload = &state->rx_buf[BADGESNAKE_HEADER_SIZE];

	switch (state->rx_buf[1]) {
	case BADGESNAKE_TOKEN_INFO:
		if (payload_len != 0U) {
			return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
								"{\"error\":\"unexpected payload\"}");
		}
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_SUCCESS,
							info_response_json);
	case BADGESNAKE_TOKEN_START:
		if (!badgesnake_payload_matches(payload, payload_len, start_request_json)) {
			return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
								"{\"error\":\"unexpected start payload\"}");
		}
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_SUCCESS, "{}");
	case BADGESNAKE_TOKEN_MOVE:
		if (!badgesnake_payload_matches(payload, payload_len, move_request_json)) {
			return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
								"{\"error\":\"unexpected move payload\"}");
		}
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_SUCCESS,
							move_response_json);
	case BADGESNAKE_TOKEN_END:
		if (!badgesnake_payload_matches(payload, payload_len, end_request_json)) {
			return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_BAD_REQUEST,
								"{\"error\":\"unexpected end payload\"}");
		}
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_SUCCESS, "{}");
	default:
		return badgesnake_prepare_json_response(state, BADGESNAKE_STATUS_UNSUPPORTED,
							"{\"error\":\"unsupported token\"}");
	}
}

static int badgesnake_write_requested(struct i2c_target_config *config)
{
	struct badgesnake_target_state *state =
		CONTAINER_OF(config, struct badgesnake_target_state, config);

	state->rx_len = 0;
	state->stop_seen = false;
	return 0;
}

static int badgesnake_write_received(struct i2c_target_config *config, uint8_t val)
{
	struct badgesnake_target_state *state =
		CONTAINER_OF(config, struct badgesnake_target_state, config);

	if (state->rx_len >= sizeof(state->rx_buf)) {
		return -ENOSPC;
	}

	state->rx_buf[state->rx_len++] = val;
	return 0;
}

static int badgesnake_read_requested(struct i2c_target_config *config, uint8_t *val)
{
	struct badgesnake_target_state *state =
		CONTAINER_OF(config, struct badgesnake_target_state, config);

	if (state->tx_pos < state->tx_len) {
		*val = state->tx_buf[state->tx_pos++];
	} else {
		*val = 0;
	}

	return 0;
}

static int badgesnake_read_processed(struct i2c_target_config *config, uint8_t *val)
{
	return badgesnake_read_requested(config, val);
}

static int badgesnake_stop(struct i2c_target_config *config)
{
	struct badgesnake_target_state *state =
		CONTAINER_OF(config, struct badgesnake_target_state, config);

	state->stop_seen = true;
	return badgesnake_dispatch_request(state);
}

static const struct i2c_target_callbacks badgesnake_callbacks = {
	.write_requested = badgesnake_write_requested,
	.write_received = badgesnake_write_received,
	.read_requested = badgesnake_read_requested,
	.read_processed = badgesnake_read_processed,
	.stop = badgesnake_stop,
};

static void badgesnake_reset_state(struct badgesnake_target_state *state)
{
	memset(state->rx_buf, 0, sizeof(state->rx_buf));
	memset(state->tx_buf, 0, sizeof(state->tx_buf));
	state->rx_len = 0;
	state->tx_len = 0;
	state->tx_pos = 0;
	state->stop_seen = false;
}

static void badgesnake_register_target_once(void)
{
	int err;

	zassert_true(device_is_ready(badgesnake_i2c), "I2C device is not ready");

	if (target_registered) {
		return;
	}

	badgesnake_reset_state(&target_state);
	target_state.config.flags = 0;
	target_state.config.address = BADGESNAKE_I2C_ADDR;
	target_state.config.callbacks = &badgesnake_callbacks;

	err = i2c_target_register(badgesnake_i2c, &target_state.config);
	zassert_ok(err, "i2c_target_register failed");
	target_registered = true;
}

static void badgesnake_expect_request_frame(const uint8_t *frame, size_t frame_len, uint8_t token,
					    const char *payload_json)
{
	size_t expected_payload_len = payload_json ? strlen(payload_json) : 0U;

	zassert_equal(frame[0], BADGESNAKE_PROTOCOL_VERSION, "unexpected request version");
	zassert_equal(frame[1], token, "unexpected request token");
	zassert_equal(frame[2], 0, "unexpected request flags");
	zassert_equal(frame[3], 0, "unexpected reserved byte");
	zassert_equal(badgesnake_u16le(&frame[4]), expected_payload_len, "unexpected payload length");
	zassert_equal(frame_len, BADGESNAKE_HEADER_SIZE + expected_payload_len,
		      "unexpected request frame length");
	if (expected_payload_len > 0U) {
		zassert_mem_equal(&frame[BADGESNAKE_HEADER_SIZE], payload_json, expected_payload_len,
				  "unexpected request payload");
	}
}

static void badgesnake_expect_response_prefix(const uint8_t *buf, size_t buf_len, uint8_t status,
					      const char *payload_json)
{
	size_t payload_len = strlen(payload_json);
	size_t frame_len = BADGESNAKE_HEADER_SIZE + payload_len;

	zassert_true(buf_len >= frame_len, "response buffer too short");
	zassert_equal(buf[0], BADGESNAKE_PROTOCOL_VERSION, "unexpected response version");
	zassert_equal(buf[1], status, "unexpected response status");
	zassert_equal(buf[2], 0, "unexpected response flags");
	zassert_equal(buf[3], 0, "unexpected response reserved byte");
	zassert_equal(badgesnake_u16le(&buf[4]), payload_len, "unexpected response payload length");
	zassert_mem_equal(&buf[BADGESNAKE_HEADER_SIZE], payload_json, payload_len,
			  "unexpected response payload");
	for (size_t i = frame_len; i < buf_len; ++i) {
		zassert_equal(buf[i], 0, "expected zero padding after response frame");
	}
}

static void badgesnake_controller_write(uint8_t token, const char *payload_json)
{
	uint8_t frame[BADGESNAKE_MAX_FRAME_SIZE];
	size_t frame_len;
	int err;

	err = badgesnake_encode_frame(token, (const uint8_t *)payload_json,
				      payload_json ? strlen(payload_json) : 0U, frame, &frame_len);
	zassert_ok(err, "failed to encode request frame");

	err = i2c_write(badgesnake_i2c, frame, frame_len, BADGESNAKE_I2C_ADDR);
	zassert_ok(err, "i2c_write failed");
}

static void badgesnake_controller_read(uint8_t *buf, size_t len)
{
	int err;

	memset(buf, 0xaa, len);
	err = i2c_read(badgesnake_i2c, buf, len, BADGESNAKE_I2C_ADDR);
	zassert_ok(err, "i2c_read failed");
}

ZTEST(badgesnake_i2c_transport, test_info_round_trip)
{
	uint8_t response[160];

	badgesnake_register_target_once();
	badgesnake_reset_state(&target_state);

	badgesnake_controller_write(BADGESNAKE_TOKEN_INFO, NULL);
	zassert_true(target_state.stop_seen, "write transaction did not complete");
	badgesnake_expect_request_frame(target_state.rx_buf, target_state.rx_len,
					BADGESNAKE_TOKEN_INFO, NULL);

	badgesnake_controller_read(response, sizeof(response));
	badgesnake_expect_response_prefix(response, sizeof(response), BADGESNAKE_STATUS_SUCCESS,
					  info_response_json);
}

ZTEST(badgesnake_i2c_transport, test_move_round_trip)
{
	uint8_t response[96];

	badgesnake_register_target_once();
	badgesnake_reset_state(&target_state);

	badgesnake_controller_write(BADGESNAKE_TOKEN_MOVE, move_request_json);
	zassert_true(target_state.stop_seen, "write transaction did not complete");
	badgesnake_expect_request_frame(target_state.rx_buf, target_state.rx_len,
					BADGESNAKE_TOKEN_MOVE, move_request_json);

	badgesnake_controller_read(response, sizeof(response));
	badgesnake_expect_response_prefix(response, sizeof(response), BADGESNAKE_STATUS_SUCCESS,
					  move_response_json);
}

ZTEST(badgesnake_i2c_transport, test_start_and_end_round_trip)
{
	uint8_t response[32];

	badgesnake_register_target_once();
	badgesnake_reset_state(&target_state);

	badgesnake_controller_write(BADGESNAKE_TOKEN_START, start_request_json);
	badgesnake_expect_request_frame(target_state.rx_buf, target_state.rx_len,
					BADGESNAKE_TOKEN_START, start_request_json);
	badgesnake_controller_read(response, sizeof(response));
	badgesnake_expect_response_prefix(response, sizeof(response), BADGESNAKE_STATUS_SUCCESS,
					  "{}");

	badgesnake_reset_state(&target_state);
	badgesnake_controller_write(BADGESNAKE_TOKEN_END, end_request_json);
	badgesnake_expect_request_frame(target_state.rx_buf, target_state.rx_len,
					BADGESNAKE_TOKEN_END, end_request_json);
	badgesnake_controller_read(response, sizeof(response));
	badgesnake_expect_response_prefix(response, sizeof(response), BADGESNAKE_STATUS_SUCCESS,
					  "{}");
}

ZTEST(badgesnake_i2c_transport, test_bad_version_is_rejected)
{
	uint8_t frame[BADGESNAKE_HEADER_SIZE];
	uint8_t response[80];
	int err;

	badgesnake_register_target_once();
	badgesnake_reset_state(&target_state);

	memset(frame, 0, sizeof(frame));
	frame[0] = 0x7f;
	frame[1] = BADGESNAKE_TOKEN_MOVE;
	err = i2c_write(badgesnake_i2c, frame, sizeof(frame), BADGESNAKE_I2C_ADDR);
	zassert_ok(err, "i2c_write failed");

	badgesnake_controller_read(response, sizeof(response));
	badgesnake_expect_response_prefix(response, sizeof(response), BADGESNAKE_STATUS_BAD_REQUEST,
					  "{\"error\":\"unsupported version\"}");
}

ZTEST_SUITE(badgesnake_i2c_transport, NULL, NULL, NULL, NULL, NULL);
