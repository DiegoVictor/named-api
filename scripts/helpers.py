import os
import pandas
import random

def read_dataset(name):
    path = os.path.dirname(__file__) + '/../datasets/'
    dataset_path = os.path.abspath(path + name + '.csv')
    return pandas.read_csv(dataset_path)

def normalize(string):
    return string.replace(
        {r'[^\w ]':''}, regex=True
    ).apply(lambda n:str(n).lower())

def random_from(items):
    rnd = random.random() * sum(items.values())
    for item in items:
        rnd -= items[item]
        if rnd < 0:
            return item

def greater_than(items, minimum):
    total = sum(items.values())
    result = dict(filter(
        lambda item: item[1] / total > minimum, items.items()
    ));

    return result