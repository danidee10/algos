"""Test for Bubble sort."""

import unittest

from bubble_sort import bubble_sort


class BubbleSortTestCase(unittest.TestCase):
    """Test case for bubble sort."""

    def test_bubble_sort(self):
        """Test bubble sort sorting algorithm."""

        unsorted_list = [14, 33, 27, 35, 10]

        sorted_list = bubble_sort(unsorted_list)

        self.assertEqual(sorted_list, [10, 14, 27, 33, 35])


if '__name__' == '__main__':
    unittest.main()
