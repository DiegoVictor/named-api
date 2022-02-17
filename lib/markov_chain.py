def build_markov_chain(dataset, word_size):
    chain = {
        '_initial':{},
        '_names': set(dataset)
    }

    for word in dataset:
        word_wrapped = str(word) + '.'

        for i in range(0, len(word_wrapped) - word_size):
            tuple = word_wrapped[i:i + word_size]
            next = word_wrapped[i + 1:i + word_size + 1]

            if tuple not in chain:
                entry = chain[tuple] = {}
            else:
                entry = chain[tuple]

            if i == 0:
                if tuple not in chain['_initial']:
                    chain['_initial'][tuple] = 1
                else:
                    chain['_initial'][tuple] += 1

            if next not in entry:
                entry[next] = 1
            else:
                entry[next] += 1
    return chain