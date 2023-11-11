import Data.Time.Clock
import Data.Time.Clock.POSIX

-- returns a random number between 1 and 100
getRandomFromTime :: IO Int
getRandomFromTime = do
    now <- getCurrentTime
    let millis = floor $ utcTimeToPOSIXSeconds now * 1000
    return $ (millis `mod` 100) + 1

status :: Int -> Int -> IO Bool
status guess secret | guess == secret = do
                                        putStrLn "You win!"
                                        return True
                    | guess > secret = do
                                       putStrLn "Too high"
                                       return False
                    | otherwise = do
                                  putStrLn "Too low"
                                  return False

play :: Int -> IO ()
play (-1) = return ()
play secret = do
    input <- getLine
    let guess = read input :: Int
    isWin <- Main.status guess secret
    if isWin then
        play (-1)
    else do
        play secret

main :: IO ()
main = do secret <- getRandomFromTime
          print secret
          putStrLn "Try to guess a secret number 1-100."
          play secret
          putStrLn "Done"
