import { test } from 'node:test';
import assert from 'node:assert/strict';
import bcrypt from 'bcryptjs';

// Test de fumée : vérifie que le hachage/verification de mot de passe fonctionne,
// sans dépendre d'une base de données (utile en CI sans Postgres).
test('bcrypt hash puis compare retourne true pour le bon mot de passe', async () => {
  const hashed = await bcrypt.hash('monMotDePasse123', 10);
  const result = await bcrypt.compare('monMotDePasse123', hashed);
  assert.equal(result, true);
});

test('bcrypt compare retourne false pour un mauvais mot de passe', async () => {
  const hashed = await bcrypt.hash('monMotDePasse123', 10);
  const result = await bcrypt.compare('mauvaisMotDePasse', hashed);
  assert.equal(result, false);
});
