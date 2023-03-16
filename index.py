
poems = '''
those who do not feel this love
pulling the like a river
those who do not drink dawn
like a cup of spring water
or take in sunset like supper
those who do not want to change
let the sleep
'''
def freq(paragraph):
    # dict frequency
    frenquency = {} 
    for word in poems.split():
        #if frenquency[word] = 0 -> frenquency.get(word, 0) = 0
        #else frenquency.get(word, 0) = frenquency[word] (e.g. 1,2,3...)
        frenquency[word] = frenquency.get(word, 0) + 1
    
    most_common = max(frenquency, key=frenquency.get)
    print(most_common)
    
def split_ext(path):
    i = path.rfind('.')
    if i == -1:
        return path, ''
    return path[:i], path[i:]
    
root, ext = split_ext('app.py')
print(f'{root=}, {ext=}')
print(f'{root}, {ext}')

freq(poems)