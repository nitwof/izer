import sys


def main():
    symbol = sys.argv[1].encode().decode("unicode-escape")
    for i in range(0, 16):
        for j in range(0, 16):
            code = str(i * 16 + j)
            sys.stdout.write(u"\u001b[38;5;" + code + "m ")
            sys.stdout.write((code + ' ' + symbol).ljust(6))
        print(u"\u001b[0m")


if __name__ == "__main__":
    main()
