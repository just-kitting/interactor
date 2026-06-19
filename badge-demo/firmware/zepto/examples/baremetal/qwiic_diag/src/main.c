#include "zepto_board.h"

int main(void) {
    zepto_board_init();
    zepto_qwiic_gpio_init();

#if defined(ZEPTO_QWIIC_DIAG_HOLD_SDA)
    zepto_qwiic_sda_drive_low();
#elif defined(ZEPTO_QWIIC_DIAG_HOLD_SCL)
    zepto_qwiic_scl_drive_low();
#else
    zepto_qwiic_release();
#endif

    while (1) {
        zepto_led_toggle();
        zepto_delay_ms(500);
    }
}
