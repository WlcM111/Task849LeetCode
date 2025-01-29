# Maximum Distance to Closest Person

This program calculates the maximum distance to the closest person (represented by `1`) in a row of seats (represented by `0`s and `1`s). It is useful for scenarios like finding the best seat in a theater or classroom where you want to maximize your distance from others.

## Function

- **maxDistToClosest(seats []int) int**: Computes the maximum distance to the closest occupied seat (`1`) in the given list of seats.

## How It Works

1. **Initialization**:
   - `start`: Tracks the starting index of a continuous block of empty seats (`0`s).
   - `res`: Stores the maximum distance found so far.
   - `cnt`: Counts the number of consecutive empty seats.

2. **Iteration**:
   - Iterates through the `seats` array.
   - If the current seat is empty (`0`):
     - Increments the `cnt` (count of consecutive empty seats).
     - If it's the first empty seat in a block, updates `start`.
     - If it's the last seat in the array, returns the maximum of `cnt` and `res`.
   - If the current seat is occupied (`1`):
     - Calculates the distance based on whether the block of empty seats is at the start, middle, or end of the row.
     - Updates `res` with the maximum distance found.
     - Resets `cnt` to 0 for the next block of empty seats.

3. **Edge Cases**:
   - Handles empty seats at the beginning or end of the row separately to ensure correct distance calculation.

4. **Return**:
   - Returns the maximum distance to the closest occupied seat.

## Example

For the input `[1, 0, 0, 0, 1, 0, 1]`, the program will return `2`, as the maximum distance to the closest person is 2 (for the seat at index `2`).
