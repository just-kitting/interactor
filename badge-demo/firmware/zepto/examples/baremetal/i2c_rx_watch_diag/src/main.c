#include "zepto_board.h"
#include "zepto_i2c_target.h"

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifndef ZEPTO_I2C_DIAG_ADDR
#define ZEPTO_I2C_DIAG_ADDR 0x42u
#endif

int main(void) {
    zepto_i2c_target_t target;
    const uint8_t *request = NULL;
    uint16_t request_len = 0u;
    bool overflowed = false;
    bool latched = false;

    zepto_board_init();
    zepto_i2c_target_init(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR);

    while (1) {
        zepto_i2c_target_poll(&target);
        if (!latched &&
            zepto_i2c_target_take_request(&target, &request, &request_len, &overflowed)) {
            if (!overflowed && request_len != 0u) {
                latched = true;
                zepto_qwiic_gpio_init();
                zepto_qwiic_scl_drive_low();
            }
        }
        zepto_led_toggle();
        zepto_delay_ms(latched ? 120u : 500u);
    }
}
