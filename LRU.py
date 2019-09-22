# coding=utf8
from collections import OrderedDict


class LRU(object):
    def __init__(self, capacity=16):
        self.capacity = capacity
        self.cache = OrderedDict()

    def get(self, key, default=None):
        if key in self.cache:
            val = self.cache.pop(key)
            self.cache[key] = val
            return val
        else:
            return default

    def set(self, key, val):
        if key in self.cache:
            self.cache.pop(key)
        if len(self.cache) >= self.capacity:
            self.cache.popitem(last=False)
        self.cache[key] = val
        return

    def __str__(self):
        return str(self.cache)


lru = LRU(2)
lru.set(1, 5)
print lru

lru.set(2, 5)
print lru

lru.set(1, 5)
print lru

lru.set(2, 7)
print lru
