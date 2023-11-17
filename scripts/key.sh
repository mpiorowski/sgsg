openssl genpkey -algorithm RSA -out private.key -pkeyopt rsa_keygen_bits:2048
openssl rsa -pubout -in private.key -out public.key

mv private.key ../client/src/lib/server/private.key
mv public.key ../server/public.key
