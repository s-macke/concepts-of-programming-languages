abs :: Int -> Int
abs n
  | n < 0     = -n
  | otherwise =  n

main = do
       print $ Main.abs 2
       print $ Main.abs (-2)
       print $ Main.abs (0)