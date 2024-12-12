# Client-Server Application

## Overview
This application is built on a client-server architecture where multiple clients communicate with a single server. The server processes data concurrently using **Goroutines** in Go.

## Key Features:
1. **Client-Server Model**: Multiple clients can interact with a single server.
2. **Concurrent Processing**: The server processes data using Goroutines for efficient concurrency.
3. **Configurable Server**: A configuration file that can be modified to customize some properties of the server defines:
   - Maximum number of elements in an array sent by clients.
   - Maximum number of active Goroutines at any time.

## Functionality
The server is capable of solving 12 distinct problems. Clients can request solutions to these problems by sending appropriate data.

## List of Problems:

### 1. **Character Extraction**  
   - **Input**: An array of strings of the same size.  
   - **Output**: An array where the `i-th` word is formed from the `i-th` character of each input word.  
   - **Example**:  
     Input: `["casa", "masa", "trei", "tanc", "4321"]`  
     Output: `["cmtt4", "aara3", "ssen2", "aaic1"]`

---

### 2. **Count Perfect Squares**  
   - **Input**: An array of strings containing mixed characters and digits.  
   - **Output**: Count of perfect square numbers.  
   - **Example**:  
     Input: `["abd4g5", "1sdf6fd", "fd2fdsf5"]`  
     Output: `2 (16 and 25)`  

---

### 3. **Sum of Reversed Numbers**  
   - **Input**: An array of integers.  
   - **Output**: Sum of the reversed numbers.  
   - **Example**:  
     Input: `[12, 13, 14]`  
     Output: `93 (21 + 31 + 41)`

---

### 4. **Arithmetic Mean in Range**  
   - **Input**: Array with interval bounds `[a, b]` as the first two elements and a series of integers.  
   - **Output**: Arithmetic mean of numbers whose digit sum falls in `[a, b]`.  
   - **Example**:  
     Input: `[2, 10, 11, 39, 32, 80, 84]`  
     Output: `41`

---

### 5. **Binary Conversion**  
   - **Input**: An array of strings.  
   - **Output**: Converts valid binary strings to base-10 integers and returns them.  
   - **Example**:  
     Input: `["2dasdas", "12", "dasdas", "1010", "101"]`  
     Output: `[10, 5]`

---

### 6. **Caesar Cipher**  
   - **Input**: Array with shift value `k` and LEFT or RIGHT as the first two elements and a series of strings.  
   - **Output**: Encoded strings using Caesar Cipher.  
   - **Example**:  
     Input: `[3, LEFT, "abcdef"]`  
     Output: `["xyzabc"]`  
     Input: `[1, RIGHT, "abcdef"]`  
     Output: `["bcdefg"]`

---

### 7. **Run-Length Decoding**  
   - **Input**: Encoded text with format `<count><character>`, where `count` is between 0 and 20.  
   - **Output**: Decoded text.  
   - **Example**:  
     Input: `"1G11o1L"`  
     Output: `"GoooooooooooL"`

---

### 8. **Prime Digit Count**  
   - **Input**: Array of natural numbers.  
   - **Output**: Total number of digits in all prime numbers.  
   - **Example**:  
     Input: `[23, 17, 15, 3, 18]`  
     Output: `5 (23, 17, 3)`

---

### 9. **Count Even-Vowel Words**  
   - **Input**: Array of strings with mixed characters and digits.  
   - **Output**: Count of words with even vowels at even positions.  
   - **Example**:  
     Input: `["mama", "iris", "bunica", "ala"]`  
     Output: `2 ("iris", "ala")`

---

### 10. **Greatest Common Divisor (GCD)**  
   - **Input**: Array of strings containing numbers and text.  
   - **Output**: GCD of valid integers.  
   - **Example**:  
     Input: `[24, 16, 8, "aaa", "bbb"]`  
     Output: `8`

---

### 11. **Right Rotate and Sum**  
   - **Input**: Array of integers that has on the first position a rotation value (`k`).  
   - **Output**: Sum of numbers after rotating digits to the right by `k` positions.  
   - **Example**:  
     Input: `[1234, 3456, 4567], k=2`  
     Output: `15791 (3412 + 5634 + 6745)`

---

### 12. **Double First Digit**  
   - **Input**: Array of natural numbers.  
   - **Output**: Sum of numbers after doubling the first digit of each.  
   - **Example**:  
     Input: `[23, 43, 26, 74]`  
     Output: `1666 (223 + 443 + 226 + 774)`
