import csv
from glob import glob
from datetime import datetime

def get_files(dir: str = 'data') -> list:
    files = glob(f'{dir}/*.csv')
    return files

def read_csv(file: str) -> iter:
    with open(file, "r") as f:
        reader = csv.reader(f)
        next(reader, None)  # skip header
        rows = [row for row in reader]
    return rows


def work_records(records: iter) -> None:
    total = 0
    for record in records:
        if 'member' in record[12]:
            total += 1
    print("the file had {v} member rides in it".format(v=str(total)))


def main():
    t1 = datetime.now()
    files = get_files()
    for file in files:
        records = read_csv(file)
        work_records(records)
    t2 = datetime.now()
    print(f"{t2}")

main()
