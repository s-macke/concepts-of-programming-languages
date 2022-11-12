signum :: Int -> Int
signum a  | a > 0 = 1
          | a == 0 = 0
          | otherwise = -1

main = print $ Main.signum -2