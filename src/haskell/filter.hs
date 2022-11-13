filter :: (a -> Bool) -> [a] -> [a]
filter p [] = []
filter p (x:xs) = if p x then x : Main.filter p xs else Main.filter p xs

main = print (Main.filter even [1..10])
