"""Bubble sort Implementation in python."""

UNSORTED_LIST = [14, 33, 27, 35, 10]


def sort(unsorted_list, swapped=False):
    """Iterate through an unsorted list and swap elements accordingly."""
    for index in range(len(unsorted_list) - 1):
        if unsorted_list[index] > unsorted_list[index + 1]:
            a, b = unsorted_list[index], unsorted_list[index + 1]
            unsorted_list[index + 1], unsorted_list[index] = a, b

            swapped = True

    return swapped

while True:
    SWAPPED = sort(UNSORTED_LIST)

    if not SWAPPED:
        break


print(UNSORTED_LIST)
