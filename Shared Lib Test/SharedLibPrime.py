import ctypes
from time import time
start = time()
so = ctypes.cdll.LoadLibrary('./preload')

prime = so.isPrime
prime.argtypes=[ctypes.c_int64]
prime.restype = ctypes.c_int64

total = so.totalPrime
total.argtypes=[ctypes.c_int64]

print(total(10000000))
print("Execution time : ", time() - start)
