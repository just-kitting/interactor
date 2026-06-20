#include "zepto_board.h"
#include "zepto_i2c_target.h"

#include <stdint.h>

#ifndef ZEPTO_I2C_DIAG_STAGE
#define ZEPTO_I2C_DIAG_STAGE ZEPTO_I2C_TARGET_STAGE_FULL
#endif

#ifndef ZEPTO_I2C_DIAG_ADDR
#define ZEPTO_I2C_DIAG_ADDR 0x42u
#endif

static uint32_t stage_delay_ms(void) {
    return 100u + ((uint32_t)ZEPTO_I2C_DIAG_STAGE * 100u);
}

int main(void) {
    zepto_i2c_target_t target;
    static const uint8_t response[] = {0x5A, 0xA5, 0x42, (uint8_t)ZEPTO_I2C_DIAG_STAGE};
    const uint8_t stage = (uint8_t)ZEPTO_I2C_DIAG_STAGE;

    zepto_board_init();
    zepto_i2c_target_init_stage(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR, stage);
    if (stage >= ZEPTO_I2C_TARGET_STAGE_ENABLE_ONLY) {
        zepto_i2c_target_set_response(&target, response, (uint16_t)sizeof(response));
    }

    while (1) {
        zepto_led_toggle();
        if (stage >= ZEPTO_I2C_TARGET_STAGE_ENABLE_ONLY) {
            zepto_i2c_target_poll(&target);
        }
        zepto_delay_ms(stage_delay_ms());
    }
}
