
# Crack hashes

VastAI Guide:
https://adamsvoboda.net/password-cracking-in-the-cloud-with-hashcat-vastai/

Open a terminal on your local machine and create a new SSH key pair and add the public key to your Vast.ai account

Rainbow tables:
https://weakpass.com/tools/lookup

Top 100k rockyou:
https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10-million-password-list-top-100000.txt

Dizionario italiano:
https://github.com/sigmasaur/AnagramSolver/blob/main/dictionary.txt

Nomi italiani:
https://gist.githubusercontent.com/pdesterlich/2562329/raw/7c09ac44d769539c61df15d5b3c441eaebb77660/nomi_italiani.txt

Copy file:

    scp -P 10029 file.txt root@ip:/root

    scp -P 10029 root@ip:/root/cracked.txt /home/username/hash/

Da dizionario:

    hashcat -m 0 -a 0 -o cracked.txt hashes.txt dictionary.txt 

Mask:

    hashcat -m 0 -a 3 hashes.txt ?d?d?d?d?d?d

Dizionario + mask:

    hashcat -m 0 -a 6 -o cracked.txt hashes.txt dictionary.txt '?d?1' -1 '!#$%&()*-.@[]^_{|}'
