import random
import sys
import getopt

from markov_chain import build_markov_chain
from helpers import read_dataset, normalize, random_from, greater_than

WORD_MIN_PERCENT_USAGE = 0.3

def generate_name(chain):
    word_part = random_from(chain['_initial'])
    result = [word_part]

    while True:
        next_parts = chain[word_part]
        above_minimum = greater_than(next_parts, WORD_MIN_PERCENT_USAGE)
        word_part = random_from(above_minimum)

        last_character = word_part[-1]
        if last_character == '.':
            break
        result.append(last_character)

    return ''.join(result)


ATTEMPTS_LIMIT = 20
ALLOW_REPEATED = False

def generate_list(amount, chain):
    attempts = 0;
    output = [];

    while len(output) < amount and attempts < ATTEMPTS_LIMIT:
        name = generate_name(chain)

        if ALLOW_REPEATED or len(chain['_names']) < 30:
            output.append(name)
        else:
            if name not in chain['_names'] and name not in output:
                attempts = 0
                output.append(name)
            else:
                attempts += 1

    return output

WORD_SIZE = 3
NAMES_AMOUNT = 6

def main(args = None):
    if len(args) == 1:
        dataset = read_dataset(args.pop())

        dataset['name'] = normalize(dataset['name'])
        chain = build_markov_chain(dataset['name'].tolist(), WORD_SIZE)
        
        print(','.join(generate_list(NAMES_AMOUNT, chain)))
    else:
        print('Usage: python main.py <dataset>')
    

main(sys.argv[1:])