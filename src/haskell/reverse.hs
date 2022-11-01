reverse :: [a] -> [a]
reverse []     = []
reverse (x:xs) = Main.reverse xs ++ [x]

main = print $ Main.reverse "Hello"