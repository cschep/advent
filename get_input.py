#!/usr/bin/env python3

import click
import requests

cookies = {
    "session": "53616c7465645f5f97efeeef0e60a7504809b2aad998053c99c4d8a598b6466c24fd9704c2e87e47f44924cba72f614c6af1698caf67af94333ccf587b60a968"
}


@click.command()
@click.option("--day", prompt="Which Day?", help="The day's input to download.")
@click.option("--year", default="2022", help="The year's input to download.")
def get(day, year):
    url = f"https://adventofcode.com/{year}/day/{day}/input"
    resp = requests.get(url, cookies=cookies)
    with open(f"{day}.input", "w") as f:
        f.write(resp.text)


if __name__ == "__main__":
    get()
