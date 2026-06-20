#include "zepto_i2c_target.h"
#include "zepto_board.h"

#include <stddef.h>
#include <stdint.h>

#ifndef ZEPTO_I2C_TARGET_USE_UNICOMM
#define ZEPTO_I2C_TARGET_USE_UNICOMM 0
#endif

#define ZEPTO_SYSCTL_BASE 0x400AF000u

#define ZEPTO_SYSCTL_SYSOSCCFG_OFFSET 0x1100u
#define ZEPTO_SYSCTL_MCLKCFG_OFFSET   0x1104u
#define ZEPTO_SYSCTL_CLKSTATUS_OFFSET 0x1204u

#define ZEPTO_SYSOSCCFG_DISABLE          (1u << 10)
#define ZEPTO_SYSOSCCFG_DISABLESTOP      (1u << 9)
#define ZEPTO_SYSOSCCFG_USE4MHZSTOP      (1u << 8)
#define ZEPTO_SYSOSCCFG_FREQ_MASK        0x3u
#define ZEPTO_MCLKCFG_USELFCLK           (1u << 20)
#define ZEPTO_MCLKCFG_MDIV_MASK          0xFu
#define ZEPTO_CLKSTATUS_CURMCLKSEL       (1u << 17)
#define ZEPTO_CLKSTATUS_SYSOSCFREQ_MASK  0x3u

#define ZEPTO_I2C0_BASE 0x400F0000u

#define ZEPTO_I2C_PWREN_OFFSET    0x0800u
#define ZEPTO_I2C_RSTCTL_OFFSET   0x0804u
#define ZEPTO_I2C_CLKCFG_OFFSET   0x0808u
#define ZEPTO_I2C_CLKDIV_OFFSET   0x1000u
#define ZEPTO_I2C_PDBGCTL_OFFSET  0x1018u
#define ZEPTO_I2C_RIS_OFFSET      0x1030u
#define ZEPTO_I2C_ICLR_OFFSET     0x1048u

#define ZEPTO_PWREN_KEY      (0x26u << 24)
#define ZEPTO_RSTCTL_KEY     (0xB1u << 24)
#define ZEPTO_CLKCFG_KEY     (0xA9u << 24)

#define ZEPTO_CLKSEL_BUSCLK  (1u << 3)
#define ZEPTO_PDBGCTL_FREE   (1u << 0)
#define ZEPTO_PDBGCTL_SOFT   (1u << 1)

#define ZEPTO_I2C_RIS_RXDONE      (1u << 16)
#define ZEPTO_I2C_RIS_TXDONE      (1u << 17)
#define ZEPTO_I2C_RIS_RXTRG       (1u << 18)
#define ZEPTO_I2C_RIS_TXTRG       (1u << 19)
#define ZEPTO_I2C_RIS_RXFULL      (1u << 20)
#define ZEPTO_I2C_RIS_TXEMPTY     (1u << 21)
#define ZEPTO_I2C_RIS_START       (1u << 22)
#define ZEPTO_I2C_RIS_STOP        (1u << 23)
#define ZEPTO_I2C_RIS_GENCALL     (1u << 24)
#define ZEPTO_I2C_RIS_DMA_DONE_TX (1u << 25)
#define ZEPTO_I2C_RIS_DMA_DONE_RX (1u << 26)
#define ZEPTO_I2C_RIS_PEC_RX_ERR  (1u << 27)
#define ZEPTO_I2C_RIS_TX_UNFL     (1u << 28)
#define ZEPTO_I2C_RIS_RX_OVFL     (1u << 29)
#define ZEPTO_I2C_RIS_ARBLOST     (1u << 30)

#define ZEPTO_I2C_CLEAR_ALL ( \
    ZEPTO_I2C_RIS_RXDONE | ZEPTO_I2C_RIS_TXDONE | ZEPTO_I2C_RIS_RXTRG | \
    ZEPTO_I2C_RIS_TXTRG | ZEPTO_I2C_RIS_RXFULL | ZEPTO_I2C_RIS_TXEMPTY | \
    ZEPTO_I2C_RIS_START | ZEPTO_I2C_RIS_STOP | ZEPTO_I2C_RIS_GENCALL | \
    ZEPTO_I2C_RIS_DMA_DONE_TX | ZEPTO_I2C_RIS_DMA_DONE_RX | ZEPTO_I2C_RIS_PEC_RX_ERR | \
    ZEPTO_I2C_RIS_TX_UNFL | ZEPTO_I2C_RIS_RX_OVFL | ZEPTO_I2C_RIS_ARBLOST)

enum {
    ZEPTO_I2C_MODE_IDLE = 0,
    ZEPTO_I2C_MODE_RX = 1,
    ZEPTO_I2C_MODE_TX = 2,
};

#if ZEPTO_I2C_TARGET_USE_UNICOMM

#define ZEPTO_I2C_CLKSEL_OFFSET   0x1008u
#define ZEPTO_I2C_IPMODE_OFFSET   0x1100u
#define ZEPTO_I2C_CTR_OFFSET      0x1100u
#define ZEPTO_I2C_ACKCTL_OFFSET   0x1104u
#define ZEPTO_I2C_SR_OFFSET       0x1108u
#define ZEPTO_I2C_IFLS_OFFSET     0x110Cu
#define ZEPTO_I2C_GFCTL_OFFSET    0x1118u
#define ZEPTO_I2C_TXDATA_OFFSET   0x1120u
#define ZEPTO_I2C_RXDATA_OFFSET   0x1124u
#define ZEPTO_I2C_OAR_OFFSET      0x114Cu

#define ZEPTO_IPMODE_I2C_TARGET       0x03u
#define ZEPTO_CTR_ENABLE              (1u << 0)
#define ZEPTO_CTR_GENCALL             (1u << 1)
#define ZEPTO_CTR_TXEMPTY_ON_TREQ     (1u << 3)
#define ZEPTO_CTR_TXWAIT_STALE_TXFIFO (1u << 5)
#define ZEPTO_CTR_RXFULL_ON_RREQ      (1u << 6)
#define ZEPTO_CTR_CLKSTRETCH          (1u << 20)
#define ZEPTO_CTR_WUEN                (1u << 21)
#define ZEPTO_OAR_OAREN               (1u << 14)

#define ZEPTO_SR_RREQ           (1u << 0)
#define ZEPTO_SR_TREQ           (1u << 1)
#define ZEPTO_SR_RXMODE         (1u << 2)
#define ZEPTO_SR_TXMODE         (1u << 7)
#define ZEPTO_SR_STALE_TXFIFO   (1u << 8)
#define ZEPTO_SR_RXCLR          (1u << 9)
#define ZEPTO_SR_TXCLR          (1u << 10)
#define ZEPTO_SR_RXFE           (1u << 11)
#define ZEPTO_SR_TXFF           (1u << 14)

#define ZEPTO_IFLS_TXCLR        (1u << 3)
#define ZEPTO_IFLS_RXCLR        (1u << 7)

#else

#define ZEPTO_I2C_CLKSEL_OFFSET    0x1004u
#define ZEPTO_I2C_GFCTL_OFFSET     0x1200u
#define ZEPTO_I2C_TOAR_OFFSET      0x1250u
#define ZEPTO_I2C_TCTR_OFFSET      0x1258u
#define ZEPTO_I2C_TSR_OFFSET       0x125Cu
#define ZEPTO_I2C_TRXDATA_OFFSET   0x1260u
#define ZEPTO_I2C_TTXDATA_OFFSET   0x1264u
#define ZEPTO_I2C_TACKCTL_OFFSET   0x1268u
#define ZEPTO_I2C_TFIFOCTL_OFFSET  0x126Cu
#define ZEPTO_I2C_TFIFOSR_OFFSET   0x1270u

#define ZEPTO_TOAR_OAREN                (1u << 14)
#define ZEPTO_TCTR_ACTIVE               (1u << 0)
#define ZEPTO_TCTR_GENCALL              (1u << 1)
#define ZEPTO_TCTR_CLKSTRETCH           (1u << 2)
#define ZEPTO_TCTR_TXEMPTY_ON_TREQ      (1u << 3)
#define ZEPTO_TCTR_TXWAIT_STALE_TXFIFO  (1u << 5)
#define ZEPTO_TCTR_WUEN                 (1u << 10)

#define ZEPTO_TSR_RREQ                  (1u << 0)
#define ZEPTO_TSR_TREQ                  (1u << 1)
#define ZEPTO_TSR_RXMODE                (1u << 2)
#define ZEPTO_TSR_TXMODE                (1u << 7)
#define ZEPTO_TSR_STALE_TXFIFO          (1u << 8)

#define ZEPTO_TFIFOCTL_TXFLUSH          (1u << 7)
#define ZEPTO_TFIFOCTL_RXFLUSH          (1u << 15)
#define ZEPTO_TFIFOCTL_RXTRIG_1_BYTE    (1u << 8)
#define ZEPTO_TFIFOCTL_TXTRIG_7_BYTES   0x7u

#define ZEPTO_TFIFOSR_RXFIFOCNT_MASK    0xFu
#define ZEPTO_TFIFOSR_RXFLUSH           (1u << 7)
#define ZEPTO_TFIFOSR_TXFIFOCNT_SHIFT   8u
#define ZEPTO_TFIFOSR_TXFIFOCNT_MASK    (0xFu << ZEPTO_TFIFOSR_TXFIFOCNT_SHIFT)
#define ZEPTO_TFIFOSR_TXFLUSH           (1u << 15)

#endif

static inline uint32_t zepto_reg_read(uint32_t addr) {
    return *(volatile uint32_t *)addr;
}

static inline void zepto_reg_write(uint32_t addr, uint32_t value) {
    *(volatile uint32_t *)addr = value;
}

static inline void zepto_i2c_write(uint32_t offset, uint32_t value) {
    zepto_reg_write(ZEPTO_I2C0_BASE + offset, value);
}

static inline uint32_t zepto_i2c_read(uint32_t offset) {
    return zepto_reg_read(ZEPTO_I2C0_BASE + offset);
}

static inline void zepto_sysctl_write(uint32_t offset, uint32_t value) {
    zepto_reg_write(ZEPTO_SYSCTL_BASE + offset, value);
}

static inline uint32_t zepto_sysctl_read(uint32_t offset) {
    return zepto_reg_read(ZEPTO_SYSCTL_BASE + offset);
}

static void zepto_i2c_target_state_init(zepto_i2c_target_t *target) {
    target->request_len = 0u;
    target->response_len = 0u;
    target->response_pos = 0u;
    target->pad_byte = 0u;
    target->transaction_mode = ZEPTO_I2C_MODE_IDLE;
    target->request_ready = false;
    target->request_overflow = false;
}

void zepto_clock_force_run_32mhz(void) {
    uint32_t sysosc_cfg = zepto_sysctl_read(ZEPTO_SYSCTL_SYSOSCCFG_OFFSET);
    uint32_t mclk_cfg = zepto_sysctl_read(ZEPTO_SYSCTL_MCLKCFG_OFFSET);

    sysosc_cfg &= ~(ZEPTO_SYSOSCCFG_DISABLE |
        ZEPTO_SYSOSCCFG_DISABLESTOP |
        ZEPTO_SYSOSCCFG_USE4MHZSTOP |
        ZEPTO_SYSOSCCFG_FREQ_MASK);
    zepto_sysctl_write(ZEPTO_SYSCTL_SYSOSCCFG_OFFSET, sysosc_cfg);

    mclk_cfg &= ~(ZEPTO_MCLKCFG_USELFCLK | ZEPTO_MCLKCFG_MDIV_MASK);
    zepto_sysctl_write(ZEPTO_SYSCTL_MCLKCFG_OFFSET, mclk_cfg);

    while ((zepto_sysctl_read(ZEPTO_SYSCTL_CLKSTATUS_OFFSET) & ZEPTO_CLKSTATUS_CURMCLKSEL) != 0u) {
    }

    while ((zepto_sysctl_read(ZEPTO_SYSCTL_CLKSTATUS_OFFSET) & ZEPTO_CLKSTATUS_SYSOSCFREQ_MASK) != 0u) {
    }
}

static void zepto_i2c_clear(uint32_t flags) {
    if (flags != 0u) {
        zepto_i2c_write(ZEPTO_I2C_ICLR_OFFSET, flags);
    }
}

#if ZEPTO_I2C_TARGET_USE_UNICOMM

static uint32_t zepto_i2c_status_read(void) {
    return zepto_i2c_read(ZEPTO_I2C_SR_OFFSET);
}

static void zepto_i2c_flush_fifos(void) {
    zepto_i2c_write(ZEPTO_I2C_IFLS_OFFSET, ZEPTO_IFLS_RXCLR | ZEPTO_IFLS_TXCLR);
    while ((zepto_i2c_status_read() & (ZEPTO_SR_RXCLR | ZEPTO_SR_TXCLR)) !=
           (ZEPTO_SR_RXCLR | ZEPTO_SR_TXCLR)) {
    }
    zepto_i2c_write(ZEPTO_I2C_IFLS_OFFSET, 0u);
}

static bool zepto_i2c_rx_not_empty(uint32_t status) {
    return (status & ZEPTO_SR_RXFE) == 0u;
}

static bool zepto_i2c_tx_can_accept(uint32_t status) {
    return (status & ZEPTO_SR_TXFF) == 0u;
}

static uint8_t zepto_i2c_read_rx_byte(void) {
    return (uint8_t)(zepto_i2c_read(ZEPTO_I2C_RXDATA_OFFSET) & 0xFFu);
}

static void zepto_i2c_write_tx_byte(uint8_t value) {
    zepto_i2c_write(ZEPTO_I2C_TXDATA_OFFSET, value);
}

static void zepto_i2c_hw_clock_init(void) {
    zepto_i2c_write(ZEPTO_I2C_IPMODE_OFFSET, ZEPTO_IPMODE_I2C_TARGET);
    zepto_i2c_write(ZEPTO_I2C_CLKDIV_OFFSET, 0u);
    zepto_i2c_write(ZEPTO_I2C_CLKSEL_OFFSET, ZEPTO_CLKSEL_BUSCLK);
    zepto_i2c_write(ZEPTO_I2C_PDBGCTL_OFFSET, ZEPTO_PDBGCTL_FREE | ZEPTO_PDBGCTL_SOFT);
}

static void zepto_i2c_hw_address_init(uint8_t address) {
    zepto_i2c_write(ZEPTO_I2C_OAR_OFFSET, ZEPTO_OAR_OAREN | (uint32_t)(address & 0x7Fu));
}

static void zepto_i2c_hw_enable_minimal(void) {
    zepto_i2c_write(ZEPTO_I2C_CTR_OFFSET, 0u);
    zepto_i2c_write(ZEPTO_I2C_ACKCTL_OFFSET, 0u);
    zepto_i2c_write(ZEPTO_I2C_GFCTL_OFFSET, 0u);
    zepto_i2c_write(
        ZEPTO_I2C_CTR_OFFSET,
        ZEPTO_CTR_ENABLE |
            ZEPTO_CTR_CLKSTRETCH |
            ZEPTO_CTR_WUEN);
}

static void zepto_i2c_hw_enable_full(void) {
    zepto_i2c_hw_enable_minimal();
    zepto_i2c_flush_fifos();
    zepto_i2c_clear(ZEPTO_I2C_CLEAR_ALL);
    zepto_i2c_write(
        ZEPTO_I2C_CTR_OFFSET,
        ZEPTO_CTR_ENABLE |
            ZEPTO_CTR_TXEMPTY_ON_TREQ |
            ZEPTO_CTR_TXWAIT_STALE_TXFIFO |
            ZEPTO_CTR_CLKSTRETCH |
            ZEPTO_CTR_WUEN);
}

#else

static uint32_t zepto_i2c_status_read(void) {
    return zepto_i2c_read(ZEPTO_I2C_TSR_OFFSET);
}

static void zepto_i2c_flush_fifos(void) {
    zepto_i2c_write(
        ZEPTO_I2C_TFIFOCTL_OFFSET,
        ZEPTO_TFIFOCTL_RXFLUSH | ZEPTO_TFIFOCTL_TXFLUSH |
            ZEPTO_TFIFOCTL_RXTRIG_1_BYTE | ZEPTO_TFIFOCTL_TXTRIG_7_BYTES);
    while ((zepto_i2c_read(ZEPTO_I2C_TFIFOSR_OFFSET) &
            (ZEPTO_TFIFOSR_RXFLUSH | ZEPTO_TFIFOSR_TXFLUSH)) !=
           (ZEPTO_TFIFOSR_RXFLUSH | ZEPTO_TFIFOSR_TXFLUSH)) {
    }
    zepto_i2c_write(
        ZEPTO_I2C_TFIFOCTL_OFFSET,
        ZEPTO_TFIFOCTL_RXTRIG_1_BYTE | ZEPTO_TFIFOCTL_TXTRIG_7_BYTES);
}

static bool zepto_i2c_rx_not_empty(uint32_t status) {
    (void)status;
    return (zepto_i2c_read(ZEPTO_I2C_TFIFOSR_OFFSET) & ZEPTO_TFIFOSR_RXFIFOCNT_MASK) != 0u;
}

static bool zepto_i2c_tx_can_accept(uint32_t status) {
    uint32_t fifo_status = zepto_i2c_read(ZEPTO_I2C_TFIFOSR_OFFSET);
    (void)status;
    return ((fifo_status & ZEPTO_TFIFOSR_TXFIFOCNT_MASK) >> ZEPTO_TFIFOSR_TXFIFOCNT_SHIFT) != 0u;
}

static uint8_t zepto_i2c_read_rx_byte(void) {
    return (uint8_t)(zepto_i2c_read(ZEPTO_I2C_TRXDATA_OFFSET) & 0xFFu);
}

static void zepto_i2c_write_tx_byte(uint8_t value) {
    zepto_i2c_write(ZEPTO_I2C_TTXDATA_OFFSET, value);
}

static void zepto_i2c_hw_clock_init(void) {
    zepto_i2c_write(ZEPTO_I2C_CLKDIV_OFFSET, 0u);
    zepto_i2c_write(ZEPTO_I2C_CLKSEL_OFFSET, ZEPTO_CLKSEL_BUSCLK);
    zepto_i2c_write(ZEPTO_I2C_PDBGCTL_OFFSET, ZEPTO_PDBGCTL_FREE | ZEPTO_PDBGCTL_SOFT);
}

static void zepto_i2c_hw_address_init(uint8_t address) {
    zepto_i2c_write(ZEPTO_I2C_TOAR_OFFSET, ZEPTO_TOAR_OAREN | (uint32_t)(address & 0x7Fu));
}

static void zepto_i2c_hw_enable_minimal(void) {
    zepto_i2c_write(ZEPTO_I2C_GFCTL_OFFSET, 0u);
    zepto_i2c_write(ZEPTO_I2C_TACKCTL_OFFSET, 0u);
    zepto_i2c_write(
        ZEPTO_I2C_TCTR_OFFSET,
        ZEPTO_TCTR_ACTIVE |
            ZEPTO_TCTR_CLKSTRETCH |
            ZEPTO_TCTR_WUEN);
}

static void zepto_i2c_hw_enable_full(void) {
    zepto_i2c_hw_enable_minimal();
    zepto_i2c_flush_fifos();
    zepto_i2c_clear(ZEPTO_I2C_CLEAR_ALL);
    zepto_i2c_write(
        ZEPTO_I2C_TCTR_OFFSET,
        ZEPTO_TCTR_ACTIVE |
            ZEPTO_TCTR_TXEMPTY_ON_TREQ |
            ZEPTO_TCTR_TXWAIT_STALE_TXFIFO |
            ZEPTO_TCTR_CLKSTRETCH |
            ZEPTO_TCTR_WUEN);
}

#endif

static void zepto_i2c_pinmux_init(void) {
    zepto_reg_write(ZEPTO_PA0_PINCM_ADDR, ZEPTO_PINCM_INENA | 0x03u);
    zepto_reg_write(ZEPTO_PA1_PINCM_ADDR, ZEPTO_PINCM_INENA | 0x03u);
}

static void zepto_i2c_pinmux_connect(void) {
    zepto_reg_write(ZEPTO_PA0_PINCM_ADDR, ZEPTO_PINCM_PC | ZEPTO_PINCM_INENA | 0x03u);
    zepto_reg_write(ZEPTO_PA1_PINCM_ADDR, ZEPTO_PINCM_PC | ZEPTO_PINCM_INENA | 0x03u);
}

static void zepto_i2c_power_init(void) {
    zepto_i2c_write(ZEPTO_I2C_PWREN_OFFSET, ZEPTO_PWREN_KEY | 1u);
    zepto_delay_cycles(128u);
    zepto_i2c_write(ZEPTO_I2C_RSTCTL_OFFSET, ZEPTO_RSTCTL_KEY | 1u);
    zepto_i2c_write(ZEPTO_I2C_CLKCFG_OFFSET, ZEPTO_CLKCFG_KEY);
    zepto_i2c_write(ZEPTO_I2C_RSTCTL_OFFSET, ZEPTO_RSTCTL_KEY | (1u << 1));
}

static void zepto_i2c_start_rx(zepto_i2c_target_t *target) {
    target->transaction_mode = ZEPTO_I2C_MODE_RX;
    target->request_len = 0u;
    target->request_ready = false;
    target->request_overflow = false;
}

static void zepto_i2c_push_rx_bytes(zepto_i2c_target_t *target) {
    while (zepto_i2c_rx_not_empty(zepto_i2c_status_read())) {
        uint8_t value = zepto_i2c_read_rx_byte();
        if (target->request_len < ZEPTO_I2C_TARGET_MAX_REQUEST) {
            target->request[target->request_len++] = value;
        } else {
            target->request_overflow = true;
        }
    }
}

static void zepto_i2c_fill_tx_fifo(zepto_i2c_target_t *target) {
    while (zepto_i2c_tx_can_accept(zepto_i2c_status_read())) {
        uint8_t value = target->pad_byte;
        if (target->response_pos < target->response_len) {
            value = target->response[target->response_pos++];
        }
        zepto_i2c_write_tx_byte(value);
    }
}

void zepto_i2c_target_init(zepto_i2c_target_t *target, uint8_t address) {
    zepto_i2c_target_init_stage(target, address, ZEPTO_I2C_TARGET_STAGE_FULL);
}

void zepto_i2c_target_init_stage(zepto_i2c_target_t *target, uint8_t address, uint8_t stage) {
    zepto_i2c_target_state_init(target);

    zepto_clock_force_run_32mhz();
    zepto_i2c_pinmux_init();
    if (stage <= ZEPTO_I2C_TARGET_STAGE_PINMUX_ONLY) {
        return;
    }

    zepto_i2c_power_init();
    if (stage <= ZEPTO_I2C_TARGET_STAGE_POWER_ONLY) {
        return;
    }

    zepto_i2c_hw_clock_init();
    if (stage <= ZEPTO_I2C_TARGET_STAGE_CLOCK_ONLY) {
        return;
    }

    zepto_i2c_hw_address_init(address);
    if (stage <= ZEPTO_I2C_TARGET_STAGE_ADDRESS_ONLY) {
        return;
    }

    zepto_i2c_hw_enable_minimal();
    if (stage <= ZEPTO_I2C_TARGET_STAGE_ENABLE_ONLY) {
        return;
    }

    zepto_i2c_pinmux_connect();
    if (stage <= ZEPTO_I2C_TARGET_STAGE_CONNECT_ONLY) {
        return;
    }

    zepto_i2c_hw_enable_full();
}

void zepto_i2c_target_poll(zepto_i2c_target_t *target) {
    uint32_t status = zepto_i2c_status_read();
    uint32_t ris = zepto_i2c_read(ZEPTO_I2C_RIS_OFFSET);

#if ZEPTO_I2C_TARGET_USE_UNICOMM
    if ((status & ZEPTO_SR_RXMODE) != 0u && target->transaction_mode != ZEPTO_I2C_MODE_RX) {
        zepto_i2c_start_rx(target);
    } else if ((status & ZEPTO_SR_TXMODE) != 0u && target->transaction_mode != ZEPTO_I2C_MODE_TX) {
        target->transaction_mode = ZEPTO_I2C_MODE_TX;
        if ((status & ZEPTO_SR_STALE_TXFIFO) != 0u) {
            zepto_i2c_flush_fifos();
        }
    }

    if ((status & ZEPTO_SR_RXMODE) != 0u ||
        (ris & (ZEPTO_I2C_RIS_RXDONE | ZEPTO_I2C_RIS_RXTRG | ZEPTO_I2C_RIS_RXFULL)) != 0u) {
#else
    if ((status & ZEPTO_TSR_RXMODE) != 0u && target->transaction_mode != ZEPTO_I2C_MODE_RX) {
        zepto_i2c_start_rx(target);
    } else if ((status & ZEPTO_TSR_TXMODE) != 0u && target->transaction_mode != ZEPTO_I2C_MODE_TX) {
        target->transaction_mode = ZEPTO_I2C_MODE_TX;
        if ((status & ZEPTO_TSR_STALE_TXFIFO) != 0u) {
            zepto_i2c_flush_fifos();
        }
    }

    if ((status & ZEPTO_TSR_RXMODE) != 0u ||
        (ris & (ZEPTO_I2C_RIS_RXDONE | ZEPTO_I2C_RIS_RXTRG | ZEPTO_I2C_RIS_RXFULL)) != 0u) {
#endif
        zepto_i2c_push_rx_bytes(target);
        zepto_i2c_clear(ris & (ZEPTO_I2C_RIS_RXDONE | ZEPTO_I2C_RIS_RXTRG | ZEPTO_I2C_RIS_RXFULL));
    }

#if ZEPTO_I2C_TARGET_USE_UNICOMM
    if ((status & (ZEPTO_SR_TXMODE | ZEPTO_SR_TREQ)) != 0u ||
#else
    if ((status & (ZEPTO_TSR_TXMODE | ZEPTO_TSR_TREQ)) != 0u ||
#endif
        (ris & (ZEPTO_I2C_RIS_TXDONE | ZEPTO_I2C_RIS_TXTRG | ZEPTO_I2C_RIS_TXEMPTY | ZEPTO_I2C_RIS_TX_UNFL)) != 0u) {
        zepto_i2c_fill_tx_fifo(target);
        zepto_i2c_clear(ris & (ZEPTO_I2C_RIS_TXDONE | ZEPTO_I2C_RIS_TXTRG | ZEPTO_I2C_RIS_TXEMPTY | ZEPTO_I2C_RIS_TX_UNFL));
    }

    if ((ris & ZEPTO_I2C_RIS_RX_OVFL) != 0u) {
        target->request_overflow = true;
        zepto_i2c_clear(ZEPTO_I2C_RIS_RX_OVFL);
    }

    if ((ris & (ZEPTO_I2C_RIS_GENCALL | ZEPTO_I2C_RIS_ARBLOST | ZEPTO_I2C_RIS_START |
                ZEPTO_I2C_RIS_PEC_RX_ERR | ZEPTO_I2C_RIS_DMA_DONE_RX | ZEPTO_I2C_RIS_DMA_DONE_TX)) != 0u) {
        zepto_i2c_clear(ris & (ZEPTO_I2C_RIS_GENCALL | ZEPTO_I2C_RIS_ARBLOST | ZEPTO_I2C_RIS_START |
                               ZEPTO_I2C_RIS_PEC_RX_ERR | ZEPTO_I2C_RIS_DMA_DONE_RX | ZEPTO_I2C_RIS_DMA_DONE_TX));
    }

    if ((ris & ZEPTO_I2C_RIS_STOP) != 0u) {
        zepto_i2c_push_rx_bytes(target);
        if (target->transaction_mode == ZEPTO_I2C_MODE_RX && target->request_len != 0u) {
            target->request_ready = true;
        }
        if (target->transaction_mode == ZEPTO_I2C_MODE_TX) {
            zepto_i2c_flush_fifos();
        }
        target->transaction_mode = ZEPTO_I2C_MODE_IDLE;
        zepto_i2c_clear(ZEPTO_I2C_RIS_STOP);
    }
}

bool zepto_i2c_target_take_request(zepto_i2c_target_t *target, const uint8_t **data, uint16_t *len, bool *overflowed) {
    if (!target->request_ready) {
        return false;
    }

    if (data != NULL) {
        *data = target->request;
    }
    if (len != NULL) {
        *len = target->request_len;
    }
    if (overflowed != NULL) {
        *overflowed = target->request_overflow;
    }

    target->request_ready = false;
    return true;
}

bool zepto_i2c_target_set_response(zepto_i2c_target_t *target, const uint8_t *data, uint16_t len) {
    uint16_t i;

    if (len > ZEPTO_I2C_TARGET_MAX_RESPONSE) {
        return false;
    }

    for (i = 0u; i < len; ++i) {
        target->response[i] = data[i];
    }

    target->response_len = len;
    target->response_pos = 0u;
    target->pad_byte = 0u;
    zepto_i2c_flush_fifos();
    return true;
}
