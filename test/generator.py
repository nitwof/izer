import json
import sys


def main():
    datafile = sys.argv[1]
    with open(sys.argv[1]) as file:
        data = json.loads(file.read())
        data['filetypes']['unknown'] = data['default']
        data['filetypes']['.unknown'] = data['default']
        for key, value in data['filetypes'].items():
            prefix = ''
            if len(sys.argv) > 2:
                if sys.argv[2] == 'golden':
                    prefix = value['symbol'] + ' '
            print(f'{prefix}{key}')
            print(f'{prefix}somedir/test/{key}')
            if key.startswith('.') and key.count('.') == 1:
                print(f'{prefix}file{key}')
                print(f'{prefix}somedir/test/file{key}')


if __name__ == "__main__":
    main()
