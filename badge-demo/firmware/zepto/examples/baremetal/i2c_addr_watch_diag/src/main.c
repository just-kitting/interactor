#include "zepto_board.h"
#include "zepto_i2c_target.h"

#include <stdbool.h>
#include <stdint.h>

#ifndef ZEPTO_I2C_DIAG_ADDR
#define ZEPTO_I2C_DIAG_ADDR 0x42u
#endif

#ifndef ZEPTO_I2C_TARGET_USE_UNICOMM
#define ZEPTO_I2C_TARGET_USE_UNICOMM 0
#endif

#define ZEPTO_I2C0_BASE 0x400F0000u

#if ZEPTO_I2C_TARGET_USE_UNICOMM
#define ZEPTO_I2C_STATUS_OFFSET 0x1108u
#define ZEPTO_I2C_ADDRMATCH_MASK 0u
#else
#define ZEPTO_I2C_STATUS_OFFSET 0x125Cu
#define ZEPTO_I2C_ADDRMATCH_MASK (0x3FFu << 9)
#endif

#define ZEPTO_I2C_STATUS_RXMODE (1u << 2)
#define ZEPTO_I2C_STATUS_TXMODE (1u << 7)
#define ZEPTO_I2C_STATUS_TREQ   (1u << 1)
#define ZEPTO_I2C_STATUS_RREQ   (1u << 0)

static inline uint32_t reg_read(uint32_t addr) {
    return *(volatile uint32_t *)addr;
}

static inline uint32_t i2c_read(uint32_t offset) {
    return reg_read(ZEPTO_I2C0_BASE + offset);
}

static bool saw_addressed_state(void) {
    uint32_t status = i2c_read(ZEPTO_I2C_STATUS_OFFSET);

    if ((status & (ZEPTO_I2C_STATUS_RXMODE |
                   ZEPTO_I2C_STATUS_TXMODE |
                   ZEPTO_I2C_STATUS_TREQ |
                   ZEPTO_I2C_STATUS_RREQ)) != 0u) {
        return true;
    }

    return (status & ZEPTO_I2C_ADDRMATCH_MASK) != 0u;
}

int main(void) {
    zepto_i2c_target_t target;
    bool latched = false;

    zepto_board_init();
    zepto_i2c_target_init(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR);

    while (1) {
        zepto_i2c_target_poll(&target);
        if (!latched && saw_addressed_state()) {
            latched = true;
            zepto_qwiic_gpio_init();
            zepto_qwiic_scl_drive_low();
        }
        zepto_led_toggle();
        zepto_delay_ms(latched ? 120u : 500u);
    }
}
