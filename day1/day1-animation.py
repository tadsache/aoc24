import math
import sys


import pygame as pg
from aocd import get_data


class Day1:
    def __init__(self, width, height, number_list):
        pg.init()

        self.width = width
        self.height = height
        self.screen = pg.display.set_mode((self.width, self.height))
        pg.display.set_caption("AOC DAY 1")

        self.background = (76, 79, 105)
        self.text = (140, 143, 161)
        self.rosewater = (220, 138, 120)
        self.saphir = (234, 118, 203)
        self.txt = (204, 208, 218)

        self.text_row_height = 90
        self.font = pg.font.SysFont("monospace", int(0.3 * self.text_row_height))

        self.text_margin = 220  # Margin for side numbers
        self.arrow_offset = 40  # Shorter arrows
        self.arrow_end_offset = 85

        self.number_list = number_list
        self.current_index = 0
        self.result = 0
        self.update_interval = 10000
        self.last_update_time = pg.time.get_ticks()

        self.is_animating = False
        self.animation_start_time = 0
        self.animation_duration = 1000  # 1 second
        self.middle_number_y = self.height // 2

    def draw_text(self, text, x, y, align="center", color=(220, 138, 120)):
        """Draw text on the screen at the specified position."""
        text_surface = self.font.render(str(text), True, color)
        text_rect = text_surface.get_rect()
        if align == "center":
            text_rect.center = (x, y)
        elif align == "left":
            text_rect.midleft = (x, y)
        elif align == "right":
            text_rect.midright = (x, y)
        self.screen.blit(text_surface, text_rect)

    def draw_arrow(self, start, end):
        """Draw an arrow between two points."""
        color = self.text
        pg.draw.line(self.screen, color, start, end, 3)
        arrow_size = 10
        angle = math.atan2(end[1] - start[1], end[0] - start[0])
        left_arrowhead = (
            end[0] - arrow_size * math.cos(angle - 0.5),
            end[1] - arrow_size * math.sin(angle - 0.5),
        )
        right_arrowhead = (
            end[0] - arrow_size * math.cos(angle + 0.5),
            end[1] - arrow_size * math.sin(angle + 0.5),
        )
        pg.draw.polygon(self.screen, color, [end, left_arrowhead, right_arrowhead])

    def render(self):
        left_number, middle_number, right_number = self.number_list[self.current_index]

        self.screen.fill(self.background)

        self.draw_text(
            "Advent of Code Day 1", self.width // 2, self.height // 2 - 150, color=self.txt
        )

        self.draw_text(left_number, self.text_margin, self.height // 2, align="left", color=self.txt)
        self.draw_text(middle_number, self.width // 2, int(self.middle_number_y))  # Animate middle number
        self.draw_text(
            right_number, self.width - self.text_margin, self.height // 2, align="right", color=self.txt
        )

        self.draw_arrow(
            (self.width // 2 - self.arrow_offset, self.height // 2),
            (self.text_margin + self.arrow_end_offset, self.height // 2),
        )
        self.draw_arrow(
            (self.width // 2 + self.arrow_offset, self.height // 2),
            (self.width - self.text_margin - self.arrow_end_offset, self.height // 2),
        )

        self.draw_text(
            f"{self.result}", self.width // 2, self.height // 2 + 80, color=self.rosewater
        )

    def render_result(self):
        self.screen.fill(self.background)

        self.draw_text(
            "Advent of Code Day 1", self.width // 2, self.height // 2 - 150, color=self.txt
        )
        self.draw_text(
            f"Result  {self.result}", self.width // 2, self.height // 2, color=self.rosewater
        )

    def animate_middle_number(self):
        if not self.is_animating:
            return

        current_time = pg.time.get_ticks()
        elapsed_time = current_time - self.animation_start_time
        if elapsed_time >= self.animation_duration:
            self.is_animating = False
            self.middle_number_y = self.height // 2 + 70
        else:
            # Interpolate the y-position
            start_y = self.height // 2
            end_y = self.height // 2 + 50
            progress = elapsed_time / self.animation_duration
            self.middle_number_y = start_y + progress * (end_y - start_y)

    def update_numbers(self):
        current_time = pg.time.get_ticks()

        if current_time - self.last_update_time >= self.update_interval:
            if self.current_index > 0:
                self.result += self.number_list[self.current_index - 1][1]

            self.current_index = (self.current_index + 1) % len(self.number_list)
            self.last_update_time = current_time

            intervals = {1: 1000, 5: 100, 15: 10, 980: 100, 997: 1000, 999: 3000}
            if self.current_index in intervals:
                self.update_interval = intervals[self.current_index]

            self.is_animating = True
            self.animation_start_time = current_time


    def run(self):
        running = True
        while running:
            for event in pg.event.get():
                if event.type == pg.QUIT:
                    running = False

            self.update_numbers()

            if self.current_index < 999:
                self.animate_middle_number()
                self.render()
            else:
                self.render_result()
            pg.display.flip()

        pg.quit()
        sys.exit()


def get_aoc_data() -> str:
    data = get_data(year=2024, day=1)
    return data


def aoc_logic(data: str):
    nums = data.split("\n")
    left_numbers = []
    right_numbers = []
    for i in nums:
        line = i.split("   ")
        left_numbers.append(line[0])
        right_numbers.append(line[1])

    left_numbers.sort()
    right_numbers.sort()
    numbers = []
    for i in range(len(left_numbers)):
        numbers.append((left_numbers[i], abs(int(left_numbers[i]) - int(right_numbers[i])), right_numbers[i]))

    return numbers


if __name__ == '__main__':
    app = Day1(width=1280, height=760, number_list=aoc_logic(get_aoc_data()))
    app.run()
