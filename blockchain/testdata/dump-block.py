#!/usr/bin/env python3

from struct import pack
import subprocess
import sys

GRS_CLI = './grsctl'
# GRS_CLI = './groestlcoin-cli'


def run(arg):
    result = subprocess.run(arg, stdout=subprocess.PIPE)
    if result.returncode:
        print(result)
        exit(1)
    return result.stdout.decode('utf-8').strip('\n')


if len(sys.argv) != 2:
    print('{} [block number]'.format(sys.argv[0]))
    exit(1)

hash = run([GRS_CLI, 'getblockhash', sys.argv[1]])
block = run([GRS_CLI, 'getblock', hash, '0'])

block = bytearray.fromhex(block)
block = b'\xf9\xbe\xb4\xd4' + pack('<i', len(block)) + block
sys.stdout.buffer.write(block)
