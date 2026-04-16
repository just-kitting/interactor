#include "zepto_board.h"

int main(void) {
    zepto_board_init();

    while (1) {
        zepto_led_toggle();
        zepto_delay_ms(250);
    }
}
