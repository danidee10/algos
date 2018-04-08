"""Bubble sort Implementation in python."""


def bubble_sort(unsorted_list):
    """Iterate through an unsorted list and swap elements accordingly."""

    # make a copy of the list
    unsorted_copy = unsorted_list[:]

    for index in range(len(unsorted_list) - 1):
        if unsorted_copy[index] > unsorted_copy[index + 1]:
            a, b = unsorted_copy[index], unsorted_copy[index + 1]
            unsorted_copy[index + 1], unsorted_copy[index] = a, b

            return bubble_sort(unsorted_copy)

    return unsorted_copy
