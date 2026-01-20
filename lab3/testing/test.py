import re
import random

TOK_RE = re.compile(r"<[^>]*>|[A-Za-z_][A-Za-z_0-9']*|[ab]")

def read_grammar(path: str):
    rules = {}
    start = None
    with open(path, "r") as f:
        for raw in f:
            line = raw.strip()

            left, right = line.split("->", 1)
            left = left.strip()
            right = right.strip()

            tokens = TOK_RE.findall(right)

            rules.setdefault(left, []).append(tuple(tokens))
            if start is None:
                start = left

    return rules, start

def parse(rules, start_symbol, word: str) -> bool:
    n = len(word)
    nonterminals = set(rules.keys())

    def is_nonterminal(x):
        return x in nonterminals

    aug = start_symbol + "'"
    while aug in nonterminals:
        aug += "'"

    chart = [set() for _ in range(n + 1)]
    chart[0].add((aug, (start_symbol,), 0, 0))

    for i in range(n + 1):
        changed = True
        while changed:
            changed = False
            for (lhs, rhs, dot, origin) in list(chart[i]):
                if dot == len(rhs):
                    for (plhs, prhs, pdot, porigin) in list(chart[origin]):
                        if pdot < len(prhs) and prhs[pdot] == lhs:
                            ns = (plhs, prhs, pdot + 1, porigin)
                            if ns not in chart[i]:
                                chart[i].add(ns)
                                changed = True
                    continue

                next_sym = rhs[dot]

                if is_nonterminal(next_sym):
                    for prod in rules.get(next_sym, []):
                        ns = (next_sym, prod, 0, i)
                        if ns not in chart[i]:
                            chart[i].add(ns)
                            changed = True
                    continue

                if i < n and next_sym == word[i]:
                    ns = (lhs, rhs, dot + 1, origin)
                    if ns not in chart[i + 1]:
                        chart[i + 1].add(ns)

    return (aug, (start_symbol,), 1, 0) in chart[n]

def add_start_S0(rules, finals):
    rules["S0"] = [(f"<0,S,{f}>",) for f in finals]
    return "S0"


path_g = "grammar.txt"
path_ll1 = "LL(1).txt"
path_lr0 = "LR(0).txt"

g_rules, g_start = read_grammar(path_g)

ll_rules, _ = read_grammar(path_ll1)
ll_start = add_start_S0(ll_rules, finals=[3, 4])

lr_rules, _ = read_grammar(path_lr0)
lr_start = add_start_S0(lr_rules, finals=[4, 5])

words = []
for _ in range(20):
    L = random.randint(0, 10)
    words.append("".join(random.choice("ab") for _ in range(L)))

all_ok = True

print("Слово\tКС\tLL1\tLR0")
for w in words:
    r1 = parse(g_rules, g_start, w)
    r2 = parse(ll_rules, ll_start, w)
    r3 = parse(lr_rules, lr_start, w)

    same = (r1 == r2 == r3)
    if not same:
        all_ok = False

    shown = w if w != "" else "ε"
    line = f"{shown}\t{r1}\t{r2}\t{r3}"
    if not same:
        line += "\t не совпало"
    print(line)

if all_ok:
    print("Результаты совпадают")
