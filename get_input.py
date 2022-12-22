#!/usr/bin/env python3

import click
import requests

cookies = {
    "session": "53616c7465645f5f3af7baccb4e286efd48e57f41282dafaa98482426475ecf0c5fa78b2fda450f6f54047f14e5b1302793fd53651fc54a301109606e3e5bf8e"
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
