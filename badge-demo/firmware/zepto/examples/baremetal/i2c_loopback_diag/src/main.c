#include "zepto_board.h"
#include "zepto_i2c_target.h"

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#ifndef ZEPTO_I2C_DIAG_ADDR
#define ZEPTO_I2C_DIAG_ADDR 0x42u
#endif

#define ZEPTO_I2C0_BASE 0x400F0000u

#define ZEPTO_I2C_PWREN_OFFSET    0x0800u
#define ZEPTO_I2C_RSTCTL_OFFSET   0x0804u
#define ZEPTO_I2C_CLKCFG_OFFSET   0x0808u
#define ZEPTO_I2C_CLKDIV_OFFSET   0x1000u
#define ZEPTO_I2C_CLKSEL_OFFSET   0x1004u
#define ZEPTO_I2C_PDBGCTL_OFFSET  0x1018u
#define ZEPTO_I2C_CSA_OFFSET      0x1210u
#define ZEPTO_I2C_CCTR_OFFSET     0x1214u
#define ZEPTO_I2C_CSR_OFFSET      0x1218u
#define ZEPTO_I2C_CRXDATA_OFFSET  0x121Cu
#define ZEPTO_I2C_CTXDATA_OFFSET  0x1220u
#define ZEPTO_I2C_CCR_OFFSET      0x1228u

#define ZEPTO_PWREN_KEY      (0x26u << 24)
#define ZEPTO_RSTCTL_KEY     (0xB1u << 24)
#define ZEPTO_CLKCFG_KEY     (0xA9u << 24)
#define ZEPTO_CLKSEL_BUSCLK  (1u << 3)
#define ZEPTO_PDBGCTL_FREE   (1u << 0)
#define ZEPTO_PDBGCTL_SOFT   (1u << 1)

#define ZEPTO_CSA_ADDR_MASK  (0x3FFu << 1)
#define ZEPTO_CSA_DIR        (1u << 0)

#define ZEPTO_CCTR_LEN_SHIFT 16u
#define ZEPTO_CCTR_STOP      (1u << 2)
#define ZEPTO_CCTR_START     (1u << 1)
#define ZEPTO_CCTR_RUN       (1u << 0)

#define ZEPTO_CSR_IDLE       (1u << 5)
#define ZEPTO_CSR_ERR        (1u << 1)
#define ZEPTO_CSR_BUSY       (1u << 0)

#define ZEPTO_CCR_LPBK       (1u << 8)
#define ZEPTO_CCR_CLKSTRETCH (1u << 2)
#define ZEPTO_CCR_ACTIVE     (1u << 0)

static inline void reg_write(uint32_t addr, uint32_t value) {
    *(volatile uint32_t *)addr = value;
}

static inline uint32_t reg_read(uint32_t addr) {
    return *(volatile uint32_t *)addr;
}

static inline void i2c_write(uint32_t offset, uint32_t value) {
    reg_write(ZEPTO_I2C0_BASE + offset, value);
}

static inline uint32_t i2c_read(uint32_t offset) {
    return reg_read(ZEPTO_I2C0_BASE + offset);
}

static void i2c_controller_init(void) {
    i2c_write(ZEPTO_I2C_PWREN_OFFSET, ZEPTO_PWREN_KEY | 1u);
    zepto_delay_cycles(128u);
    i2c_write(ZEPTO_I2C_RSTCTL_OFFSET, ZEPTO_RSTCTL_KEY | 1u);
    i2c_write(ZEPTO_I2C_CLKCFG_OFFSET, ZEPTO_CLKCFG_KEY);
    i2c_write(ZEPTO_I2C_RSTCTL_OFFSET, ZEPTO_RSTCTL_KEY | (1u << 1));
    i2c_write(ZEPTO_I2C_CLKDIV_OFFSET, 0u);
    i2c_write(ZEPTO_I2C_CLKSEL_OFFSET, ZEPTO_CLKSEL_BUSCLK);
    i2c_write(ZEPTO_I2C_PDBGCTL_OFFSET, ZEPTO_PDBGCTL_FREE | ZEPTO_PDBGCTL_SOFT);
    i2c_write(ZEPTO_I2C_CCR_OFFSET, ZEPTO_CCR_LPBK | ZEPTO_CCR_CLKSTRETCH | ZEPTO_CCR_ACTIVE);
}

static bool i2c_wait_idle(zepto_i2c_target_t *target, uint32_t iterations) {
    while (iterations-- != 0u) {
        zepto_i2c_target_poll(target);
        if ((i2c_read(ZEPTO_I2C_CSR_OFFSET) & ZEPTO_CSR_BUSY) == 0u &&
            (i2c_read(ZEPTO_I2C_CSR_OFFSET) & ZEPTO_CSR_IDLE) != 0u) {
            return true;
        }
    }
    return false;
}

static bool i2c_send_write(zepto_i2c_target_t *target, uint8_t address, const uint8_t *data, uint32_t len) {
    uint32_t i;

    i2c_write(ZEPTO_I2C_CSA_OFFSET, ((uint32_t)(address & 0x7Fu) << 1) & ZEPTO_CSA_ADDR_MASK);
    for (i = 0u; i < len; ++i) {
        i2c_write(ZEPTO_I2C_CTXDATA_OFFSET, data[i]);
    }
    i2c_write(
        ZEPTO_I2C_CCTR_OFFSET,
        ((len & 0x0FFFu) << ZEPTO_CCTR_LEN_SHIFT) |
            ZEPTO_CCTR_STOP |
            ZEPTO_CCTR_START |
            ZEPTO_CCTR_RUN);

    if (!i2c_wait_idle(target, 200000u)) {
        return false;
    }

    return (i2c_read(ZEPTO_I2C_CSR_OFFSET) & ZEPTO_CSR_ERR) == 0u;
}

static bool i2c_send_read(zepto_i2c_target_t *target, uint8_t address, uint8_t *value) {
    i2c_write(
        ZEPTO_I2C_CSA_OFFSET,
        ((((uint32_t)address) & 0x7Fu) << 1) |
            ZEPTO_CSA_DIR);
    i2c_write(
        ZEPTO_I2C_CCTR_OFFSET,
        (1u << ZEPTO_CCTR_LEN_SHIFT) |
            ZEPTO_CCTR_STOP |
            ZEPTO_CCTR_START |
            ZEPTO_CCTR_RUN);

    if (!i2c_wait_idle(target, 200000u)) {
        return false;
    }
    if ((i2c_read(ZEPTO_I2C_CSR_OFFSET) & ZEPTO_CSR_ERR) != 0u) {
        return false;
    }

    *value = (uint8_t)(i2c_read(ZEPTO_I2C_CRXDATA_OFFSET) & 0xFFu);
    return true;
}

static bool bytes_match(const uint8_t *lhs, const uint8_t *rhs, size_t len) {
    size_t i;

    for (i = 0; i < len; ++i) {
        if (lhs[i] != rhs[i]) {
            return false;
        }
    }
    return true;
}

int main(void) {
    static const uint8_t request[] = {0x10, 0x20, 0x30};
    static const uint8_t response[] = {0xA5};
    zepto_i2c_target_t target;
    const uint8_t *rx_data = NULL;
    uint16_t rx_len = 0u;
    bool overflowed = false;
    uint8_t rx_value = 0u;
    bool ok = false;

    zepto_board_init();
    zepto_i2c_target_init(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR);
    zepto_i2c_target_set_response(&target, response, (uint16_t)sizeof(response));
    i2c_controller_init();

    ok = i2c_send_write(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR, request, (uint32_t)sizeof(request));
    if (ok) {
        ok = zepto_i2c_target_take_request(&target, &rx_data, &rx_len, &overflowed);
    }
    if (ok) {
        ok = !overflowed &&
            rx_len == (uint16_t)sizeof(request) &&
            bytes_match(rx_data, request, sizeof(request));
    }
    if (ok) {
        ok = i2c_send_read(&target, (uint8_t)ZEPTO_I2C_DIAG_ADDR, &rx_value);
    }
    if (ok) {
        ok = rx_value == response[0];
    }

    while (1) {
        zepto_led_toggle();
        zepto_delay_ms(ok ? 100u : 600u);
    }
}
