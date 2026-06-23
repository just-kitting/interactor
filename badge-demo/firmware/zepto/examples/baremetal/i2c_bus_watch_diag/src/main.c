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
#define ZEPTO_I2C_RIS_OFFSET 0x1030u

#if ZEPTO_I2C_TARGET_USE_UNICOMM
#define ZEPTO_I2C_STATUS_OFFSET 0x1108u
#define ZEPTO_I2C_STATUS_BUS_ACTIVITY_MASK ((1u << 2) | (1u << 7) | (1u << 1) | (1u << 0))
#else
#define ZEPTO_I2C_STATUS_OFFSET 0x125Cu
#define ZEPTO_I2C_STATUS_BUS_ACTIVITY_MASK ((1u << 6) | (1u << 2) | (1u << 7) | (1u << 1) | (1u << 0))
#endif

#define ZEPTO_I2C_RIS_RXDONE  (1u << 16)
#define ZEPTO_I2C_RIS_TXDONE  (1u << 17)
#define ZEPTO_I2C_RIS_RXTRG   (1u << 18)
#define ZEPTO_I2C_RIS_TXTRG   (1u << 19)
#define ZEPTO_I2C_RIS_RXFULL  (1u << 20)
#define ZEPTO_I2C_RIS_TXEMPTY (1u << 21)
#define ZEPTO_I2C_RIS_START   (1u << 22)
#define ZEPTO_I2C_RIS_STOP    (1u << 23)

static inline uint32_t reg_read(uint32_t addr) {
    return *(volatile uint32_t *)addr;
}

static inline uint32_t i2c_read(uint32_t offset) {
    return reg_read(ZEPTO_I2C0_BASE + offset);
}

static bool saw_external_activity(void) {
    uint32_t status = i2c_read(ZEPTO_I2C_STATUS_OFFSET);
    uint32_t ris = i2c_read(ZEPTO_I2C_RIS_OFFSET);

    if ((status & ZEPTO_I2C_STATUS_BUS_ACTIVITY_MASK) != 0u) {
        return true;
    }

    return (ris & (ZEPTO_I2C_RIS_RXDONE |
                   ZEPTO_I2C_RIS_TXDONE |
                   ZEPTO_I2C_RIS_RXTRG |
                   ZEPTO_I2C_RIS_TXTRG |
                   ZEPTO_I2C_RIS_RXFULL |
                   ZEPTO_I2C_RIS_TXEMPTY |
                   ZEPTO_I2C_RIS_START |
                   ZEPTO_I2C_RIS_STOP)) != 0u;
}

int main(void) {
    zepto_i2c_target_t target;
    bool latched = false;

    zepto_board_init();
    zepto_i2c_target_init(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR);

    while (1) {
        zepto_i2c_target_poll(&target);
        if (!latched && saw_external_activity()) {
            latched = true;
            zepto_qwiic_gpio_init();
            zepto_qwiic_scl_drive_low();
        }
        zepto_led_toggle();
        zepto_delay_ms(latched ? 120u : 500u);
    }
}
