#!/usr/bin/env -S uv run

import sys

def first(lines):
    dial = 50
    password = 0

    for line in lines:
        delta = int(line[1:])
        if line[0] == "L":
            delta *= -1
        dial = (dial + delta) % 100
        if dial == 0:
           password += 1

    print(f"first: {password}")

def second(lines):
    dial = 50
    password = 0

    for line in lines:
        delta = int(line[1:])
        if line[0] == "L":
            delta *= -1

        rev = dial + delta
        if rev < 0 or 99 < rev:
            password += abs(rev // 100)

        dial = rev % 100

    print(f"second: {password}")


with open("data.txt", "r") as f:
    lines = f.readlines()


first(lines)
second(lines)
