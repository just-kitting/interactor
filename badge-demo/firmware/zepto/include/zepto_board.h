#ifndef ZEPTO_BOARD_H
#define ZEPTO_BOARD_H

#include <stdint.h>

#define ZEPTO_GPIOA_BASE 0x400A0000u
#define ZEPTO_IOMUX_BASE 0x40428000u

#define ZEPTO_GPIO_DOUTSET31_0 0x1290u
#define ZEPTO_GPIO_DOUTCLR31_0 0x12A0u
#define ZEPTO_GPIO_DOUTTGL31_0 0x12B0u
#define ZEPTO_GPIO_DOESET31_0  0x12D0u

#define ZEPTO_PA12_PINCM_ADDR 0x40428084u
#define ZEPTO_PA12_BIT 12u

#define ZEPTO_PINCM_PF_GPIO 0x01u
#define ZEPTO_PINCM_PC (1u << 7)
#define ZEPTO_PINCM_INENA (1u << 18)

void zepto_board_init(void);
void zepto_led_on(void);
void zepto_led_off(void);
void zepto_led_toggle(void);
void zepto_delay_cycles(volatile uint32_t cycles);
void zepto_delay_ms(uint32_t ms);

#endif
