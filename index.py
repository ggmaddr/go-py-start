
def freq(paragraph):
    # dict frequency
    frenquency = {} 
    for word in poems.split():
        frenquency[word] = frenquency.get(word, 0) + 1

    most_common = max(frenquency, key=frenquency.get)
    print(most_common)

poems = '''
those who do not feel this love
pulling the mlike a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let the sleep
'''
freq(poems)