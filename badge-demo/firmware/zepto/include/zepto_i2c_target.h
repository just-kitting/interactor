#ifndef ZEPTO_I2C_TARGET_H
#define ZEPTO_I2C_TARGET_H

#include <stdbool.h>
#include <stdint.h>

#define ZEPTO_I2C_TARGET_MAX_REQUEST 1536u
#define ZEPTO_I2C_TARGET_MAX_RESPONSE 512u

typedef struct {
    uint8_t request[ZEPTO_I2C_TARGET_MAX_REQUEST];
    uint8_t response[ZEPTO_I2C_TARGET_MAX_RESPONSE];
    uint16_t request_len;
    uint16_t response_len;
    uint16_t response_pos;
    uint8_t pad_byte;
    uint8_t transaction_mode;
    bool request_ready;
    bool request_overflow;
} zepto_i2c_target_t;

enum {
    ZEPTO_I2C_TARGET_STAGE_PINMUX_ONLY = 0,
    ZEPTO_I2C_TARGET_STAGE_POWER_ONLY = 1,
    ZEPTO_I2C_TARGET_STAGE_CLOCK_ONLY = 2,
    ZEPTO_I2C_TARGET_STAGE_ADDRESS_ONLY = 3,
    ZEPTO_I2C_TARGET_STAGE_ENABLE_ONLY = 4,
    ZEPTO_I2C_TARGET_STAGE_CONNECT_ONLY = 5,
    ZEPTO_I2C_TARGET_STAGE_FULL = 6,
};

void zepto_i2c_target_init(zepto_i2c_target_t *target, uint8_t address);
void zepto_i2c_target_init_stage(zepto_i2c_target_t *target, uint8_t address, uint8_t stage);
void zepto_i2c_target_poll(zepto_i2c_target_t *target);
bool zepto_i2c_target_take_request(zepto_i2c_target_t *target, const uint8_t **data, uint16_t *len, bool *overflowed);
bool zepto_i2c_target_set_response(zepto_i2c_target_t *target, const uint8_t *data, uint16_t len);

#endif
