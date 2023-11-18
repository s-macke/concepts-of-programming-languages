-- - Implement a converter function of an integer to a binary string.
--
-- ```Haskell
-- intToBinary :: Int -> String
-- ....
--
--  main = print $ intTooBinary 42  -- returns "101010"
-- ```

intToBinary :: Int -> String
intToBinary 0 = "0"
intToBinary n = reverse (binary n)
  where
    binary 0 = ""
    binary n = show (n `mod` 2) ++ binary (n `div` 2)

main = do
  print $ intToBinary 5