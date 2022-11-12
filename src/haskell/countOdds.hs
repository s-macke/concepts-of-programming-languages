countOdds :: [Int] -> Int
countOdds [] = 0
countOdds (x:xs)
  | odd x     = 1 + countOdds xs
  | otherwise = countOdds xs
main = print $ countOdds [1,2,3,4,5]  -- returns 3

