CREATE TABLE "branch" (
  "id" UUID PRIMARY KEY NOT NULL,
  "name" VARCHAR(50),
  "address" VARCHAR(65),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp,
  "deleted" BOOLEAN DEFAULT false,
  "deleted_at" timestamp
);

CREATE TABLE "staff_tarif" (
  "id" UUID PRIMARY KEY NOT NULL,
  "name" VARCHAR(35) NOT NULL,
  "type" VARCHAR(35) NOT NULL DEFAULT 'fixed',
  "amount_for_cash" NUMERIC NOT NULL,
  "amount_for_card" NUMERIC NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted" BOOLEAN DEFAULT false,
  "deleted_at" TIMESTAMP

);
CREATE TABLE "staff" (
  "id" UUID PRIMARY KEY NOT NULL,
  "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
  "tarif_id" UUID NOT NULL REFERENCES "staff_tarif"("id"),
  "type" VARCHAR NOT NULL,
  "name" VARCHAR NOT NULL,
  "balance" NUMERIC NOT NULL DEFAULT 0,
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp,
  "deleted" BOOLEAN DEFAULT false,
  "deleted_at" timestamp
);

CREATE TABLE "sales" (
  "id" UUID PRIMARY KEY NOT NULL,
  "branch_id" UUID NOT NULL REFERENCES "branch"("id"),
  "shop_assistent_id" UUID REFERENCES "staff"("id"),
  "cashier_id" UUID NOT NULL REFERENCES "staff"("id"),
  "price" NUMERIC NOT NULL,
  "payment_type" VARCHAR NOT NULL,
  "client_name" VARCHAR,
  "status" VARCHAR NOT NULL DEFAULT 'success',
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "updated_at" TIMESTAMP,
  "deleted" BOOLEAN DEFAULT false,
  "deleted_at" TIMESTAMP

);

CREATE TABLE "staff_transaction" (
  "id" UUID PRIMARY KEY NOT NULL,
  "sales_id" UUID NOT NULL REFERENCES "sales"("id"),
  "type" VARCHAR NOT NULL,
  "text" TEXT NOT NULL,
  "amount" NUMERIC NOT NULL,
  "staff_id" UUID NOT NULL REFERENCES "staff"("id"),
  "created_at" timestamp DEFAULT (CURRENT_TIMESTAMP),
  "updated_at" timestamp,
  "deleted" BOOLEAN DEFAULT false,
  "deleted_at" timestamp
);



SELECT 
    COUNT(*) OVER()
FROM staff_transaction AS st
JOIN  staff AS s ON st.staff_id = s.id
JOIN sales AS sl ON sl.id = st.sales_id
WHERE sl.deleted = false AND sl.status='success' AND s.created_at = NOW()-1;