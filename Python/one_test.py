import unittest
from Python.one import solution


class TestOne(unittest.TestCase):
    def test_1000(self):
        assert (233168 == solution(1000))

    def test_10(self):
        assert (23 == solution(10))
