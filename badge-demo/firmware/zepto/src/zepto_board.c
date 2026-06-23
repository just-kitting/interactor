#include "zepto_board.h"

static inline void zepto_reg_write(uint32_t addr, uint32_t value) {
    *(volatile uint32_t *)addr = value;
}

static inline void zepto_gpio_set_output_low(uint32_t bit) {
    zepto_reg_write(ZEPTO_GPIOA_BASE + ZEPTO_GPIO_DOUTCLR31_0, 1u << bit);
    zepto_reg_write(ZEPTO_GPIOA_BASE + ZEPTO_GPIO_DOESET31_0, 1u << bit);
}

void zepto_board_init(void) {
    zepto_reg_write(ZEPTO_PA13_PINCM_ADDR, ZEPTO_PINCM_PC | ZEPTO_PINCM_PF_GPIO);
    zepto_gpio_set_output_low(ZEPTO_PA13_BIT);
}

void zepto_led_on(void) {
    zepto_reg_write(ZEPTO_GPIOA_BASE + ZEPTO_GPIO_DOUTSET31_0, 1u << ZEPTO_PA13_BIT);
}

void zepto_led_off(void) {
    zepto_reg_write(ZEPTO_GPIOA_BASE + ZEPTO_GPIO_DOUTCLR31_0, 1u << ZEPTO_PA13_BIT);
}

void zepto_led_toggle(void) {
    zepto_reg_write(ZEPTO_GPIOA_BASE + ZEPTO_GPIO_DOUTTGL31_0, 1u << ZEPTO_PA13_BIT);
}

void zepto_qwiic_gpio_init(void) {
    zepto_reg_write(ZEPTO_PA0_PINCM_ADDR, ZEPTO_PINCM_PC | ZEPTO_PINCM_INENA | ZEPTO_PINCM_PF_GPIO);
    zepto_reg_write(ZEPTO_PA1_PINCM_ADDR, ZEPTO_PINCM_PC | ZEPTO_PINCM_INENA | ZEPTO_PINCM_PF_GPIO);
    zepto_qwiic_release();
}

void zepto_qwiic_sda_drive_low(void) {
    zepto_gpio_set_output_low(ZEPTO_PA0_BIT);
}

void zepto_qwiic_scl_drive_low(void) {
    zepto_gpio_set_output_low(ZEPTO_PA1_BIT);
}

void zepto_qwiic_release(void) {
    zepto_reg_write(ZEPTO_GPIOA_BASE + ZEPTO_GPIO_DOECLR31_0, (1u << ZEPTO_PA0_BIT) | (1u << ZEPTO_PA1_BIT));
}

void zepto_delay_cycles(volatile uint32_t cycles) {
    while (cycles-- != 0u) {
        __asm__ volatile ("nop");
    }
}

void zepto_delay_ms(uint32_t ms) {
    while (ms-- != 0u) {
        zepto_delay_cycles(ZEPTO_CPU_HZ / 10000u);
    }
}
