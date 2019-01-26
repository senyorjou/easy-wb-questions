# Match correct open/closed tokens on a string
def match_token(s):
    stack = []
    for token in s:
        if token in '()':
            if token == '(':
                stack.append(')')
            elif not stack or token != stack.pop():
                return False
    return not stack

print match_token('()((()()))')
print match_token('(([)))*((()]([])x)x)')
