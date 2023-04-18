import hashlib

class Node:
    def __init__(self, key, value):
        self.key = key
        self.value = value
        self.next = None

class HashMap:
    def __init__(self):
        self.size = 16
        self.buckets = [None] * self.size

    def get_hash(self, key):
        if isinstance(key, str):
            return int(hashlib.md5(key.encode('utf-8')).hexdigest(), 16)
        else:
            return hash(key)

    def get_bucket_index(self, key):
        return self.get_hash(key) % self.size

    def put(self, key, value):
        bucket_index = self.get_bucket_index(key)
        node = self.buckets[bucket_index]

        while node is not None:
            if node.key == key:
                node.value = value
                return
            node = node.next

        new_node = Node(key, value)
        new_node.next = self.buckets[bucket_index]
        self.buckets[bucket_index] = new_node

    def get(self, key):
        bucket_index = self.get_bucket_index(key)
        node = self.buckets[bucket_index]

        while node is not None:
            if node.key == key:
                return node.value
            node = node.next

        return None
