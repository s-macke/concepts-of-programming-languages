  isPalindrome [] = True
  isPalindrome [x] = True
  isPalindrome (x:xs) = x == last xs && isPalindrome (init xs)
  main = print $ isPalindrome [1,2,3,2,1]
