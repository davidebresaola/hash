
  

# Crack Hashes Guide

Questa guida contiene le informazioni necessarie per impostare e condurre attacchi di password cracking in modo efficiente.
  

## VastAI Guide

Per utilizzare Vast.ai per il cracking delle password, fai riferimento alla guida qui:

[VastAI Guide - Password Cracking in the Cloud with Hashcat](https://adamsvoboda.net/password-cracking-in-the-cloud-with-hashcat-vastai/)

  

1. Apri un terminale sulla tua macchina locale.

2. Crea una nuova coppia di chiavi SSH e aggiungi la chiave pubblica al tuo account Vast.ai.

  

## Risorse Utili

  

-  **Rainbow Tables**: Usa rainbow tables per verificare rapidamente hash comuni.

- [Rainbow Tables su Weakpass](https://weakpass.com/tools/lookup)

  

-  **Top 100k Passwords (RockYou)**: Questo dizionario contiene le password più comuni, utile per attacchi di dizionario.

- [Top 100k RockYou Password List](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10-million-password-list-top-100000.txt)

  

-  **Dizionario Italiano**: Un elenco di parole italiane, utile per attacchi mirati con dizionari linguistici.

- [Dizionario Italiano](https://github.com/sigmasaur/AnagramSolver/blob/main/dictionary.txt)

  

-  **Nomi Italiani**: Un elenco di nomi italiani, utile per attacchi di dizionario su target italiani.

- [Nomi Italiani](https://gist.githubusercontent.com/pdesterlich/2562329/raw/7c09ac44d769539c61df15d5b3c441eaebb77660/nomi_italiani.txt)

  

## Copia File

  

Per copiare file da e verso la macchina su Vast.ai, utilizza il comando `scp` con il parametro della porta specifica di Vast.ai (esempio: `10029`).

  

**Carica un file:**

```bash

scp  -P  10029  file.txt  root@<IP>:/root

```

  

**Scarica un file:**

```bash

scp  -P  10029  root@<IP>:/root/cracked.txt  /home/username/hash/

```

  

## Comandi Hashcat

  

### Crack con Dizionario

  

Utilizza un dizionario per crackare gli hash:

```bash

hashcat  -m  0  -a  0  -o  cracked.txt  hashes.txt  dictionary.txt

```

  

### Crack con Maschera (Mask Attack)

  

Per un attacco a maschera, sostituisci i caratteri con pattern specifici come segue:

```bash

hashcat  -m  0  -a  3  hashes.txt  ?d?d?d?d?d?d

```

  

### Crack con Dizionario + Maschera

  

Utilizza un dizionario combinato con una maschera per migliorare il tasso di successo:

```bash

hashcat  -m  0  -a  6  -o  cracked.txt  hashes.txt  dictionary.txt  '?d?1'  -1  '!#$%&()*-.@[]^_{|}'

```

  

## Maschere Hashcat

  

Di seguito un elenco di maschere disponibili in Hashcat per personalizzare gli attacchi:

  

-  `?d` - Qualsiasi cifra (0–9)

-  `?l` - Lettera minuscola (a–z)

-  `?u` - Lettera maiuscola (A–Z)

-  `?s` - Qualsiasi carattere speciale (`! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ { | } ~`)

-  `?a` - Qualsiasi carattere alfanumerico (`?l, ?u, ?d, ?s`)

-  `?b` - Qualsiasi carattere ASCII (0x00 - 0xff)

  

## Cracking Hashes Using John the Ripper

  

John the Ripper è un eccellente strumento per eseguire attacchi incrementali (brute-force). Ecco come utilizzarlo in vari tipi di attacco.

  

### Installazione di John the Ripper

  

Se non è già installato sulla tua istanza Vast.ai, puoi installarlo eseguendo:

  

```bash

apt  update && apt  install  -y  john

```

  

### Esecuzione di un Attacco Incrementale

  

In modalità incrementale, John the Ripper prova ogni possibile combinazione di caratteri basata sul set specificato. Comando base:

  

```bash

john  --incremental  hashes.txt

```

  

-  **Opzioni**:

-  `--incremental`: Abilita la modalità incrementale.

-  `hashes.txt`: Il file contenente gli hash da crackare.

  

### Personalizzare l'Attacco Incrementale

  

Per personalizzare l'attacco incrementale, puoi specificare parametri come lunghezza dei caratteri, set di caratteri o regole personalizzate.

  

-  **Limitare la Lunghezza della Password**:

```bash

john --incremental --max-length=8  hashes.txt

```

  

-  **Specificare il Set di Caratteri**:

  

John the Ripper include diversi set di caratteri predefiniti (come `alpha`, `digits`, ecc.). Ecco come utilizzarli:

```bash

john --incremental=digits  hashes.txt

```

  

Oppure, definisci un set di caratteri personalizzato nel file di configurazione `john.conf` se desideri caratteri specifici, ad esempio:

  

```ini

[Incremental:Custom]

File = $JOHN/john.conf

MinLen = 1

MaxLen = 8

CharCount = 36

```

  

Quindi esegui:

```bash

john --incremental=Custom  hashes.txt

```

  

### Combinare Attacco a Dizionario e Incrementale

  

Esegui prima un attacco a dizionario per tentare password più facili, seguito da un attacco incrementale per pattern più complessi o sconosciuti:

  

-  **Attacco a Dizionario**:

```bash

john --wordlist=dictionary.txt  hashes.txt

```

  

-  **Attacco Incrementale (se il dizionario fallisce)**:

```bash

john --incremental hashes.txt

```

  

## Monitoraggio e Recupero dei Risultati

  

-  **Visualizza Password Craccate**:

```bash

john --show hashes.txt

```

  

-  **Recupera il File delle Password Craccate**:

  

Trasferisci il file `john.pot` (dove John the Ripper salva le password craccate) sulla tua macchina locale:

  

```bash

scp -P 10029 root@<IP>:/root/.john/john.pot /home/username/hash/

```

  

---

  

## Note Aggiuntive

  

-  **Modello Hashcat** (`-m`): Sostituisci `-m 0` con il tipo di hash appropriato. Ad esempio, `-m 1000` per NTLM, `-m 1800` per SHA-512.

-  **Tipo di Attacco Hashcat** (`-a`):

-  `-a 0`: Attacco a dizionario

-  `-a 3`: Attacco a maschera

-  `-a 6`: Combinazione di dizionario + maschera

-  `-a 7`: Combinazione di maschera + dizionario

  

---
