# Order

- J'ai structuré mon code en différentes parties: `adapter` pour les interfaces aux données, `controller` pour les API, `model` pour mes objets métiers. Je pourrai rajouter `usecase` pour mieux isoler la partie purement objet et le reste de l'exécution.
- J'ai une seule fonction d'écriture en base pour pouvoir faire une écriture transactionnelle. Cela me semble pertinent pour garder une cohérence dans les données: Un même `product` peut apparaître dans plusieurs `order` et ne peut pas ne pas appartenir à au moins 1 `order`. Un `order` à au moins 1 `product`.
- Pour assumer un burst, j'ai implémenté un chanel simple sans limite pour augmenter la disponibilité. Le handler lit le ticket et l'envoie dans le chanel. n workers depilent ce chanel pour update la base de donnée. Si un ticket n'arrive pas à être traité je le print simplement dans la console. Ces tickets en erreurs pourraient être stockés pour aller jusqu'au bout de la consigne de n'en perdre aucun. 

Voici ma stucture de donnée pour `orderdb`:

![](https://github.com/haagor/orderMP/blob/main/db.png)

Quelques points mis de côté :
- tests
- utiliser le context depuis le handler et le descendre dans l'adapter
- remplacer les prints par du log
- passer la conf de la db dans des variables d'environnement

---
Simon P