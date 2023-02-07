import time
print("this program needs to be ran in a newer console such as windows terminal")
time.sleep(5)
#color vars
red = lambda text: '\033[0;31m' + text + '\033[0m'
green = lambda text: '\033[0;32m' + text + '\033[0m'
yellow = lambda text: '\033[0;33m' + text + '\033[0m'
blue = lambda text: '\033[0;34m' + text + '\033[0m'
magenta = lambda text: '\033[0;35m' + text + '\033[0m'

continueYN = 'Y'
 
while continueYN == 'Y' or 'y':
   sDegreeF = input('Enter temperature in degrees Fahrenheit (F):')
 
   #make var nDegreeF and int
   nDegreeF = int(sDegreeF)
 
   #check temp
   if nDegreeF < 32:
      print(blue('Pack long underwear and wear good Winter gear!'))
   elif nDegreeF > 32 and nDegreeF < 55:
      print(green('Wear a Coat!'))
   elif nDegreeF > 55 and nDegreeF < 65:
      print(green('You don\'t need a coat, but at least wear a thin sweatshirt!'))
   elif nDegreeF > 65 and nDegreeF < 85:
      print(yellow('Put on some sunscreen!'))
   elif nDegreeF > 85 and nDegreeF < 100:
      print(yellow('It\'s getting a little hot outside, maybe it is time to go inside?'))
   elif nDegreeF > 100 and nDegreeF < 120:
      print(red('Get inside!!! It\'s getting a very dangerous!!!'))
   elif nDegreeF > 120 and nDegreeF < 120000000000000000:
      print(magenta('Get inside!!!!!!!!!!!!!!!!!!!!!!!'))
 
   continueYN = input('Input another? Y/N \n')
