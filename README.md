# 🧙‍♂️ Gandalf Detector – 2FA Wallet Security for Venn Network

> *"You shall not pass... without 2FA."* – Gandalf (probably)

**Gandalf Detector** is a security-first transaction detection and authorization tool for wallets on the Venn Network. It integrates Two-Factor Authentication (2FA) using **Discord**, enabling users to add strong protection to sensitive operations like high-value transfers or contract interactions.

---

## 🚀 Overview

The **Gandalf Detector** lets users enforce 2FA on their Venn-based wallets by requiring confirmation via **Discord** before executing high-risk transactions.

Powered by a smart contract for on-chain enrollment and paired with a Discord bot, Gandalf acts as the wise gatekeeper of your wallet—blocking unauthorized or risky operations with multi-layered verification.

---

## 🔗 Key Components

- 🧾 **2FA Enrollment Smart Contract and DApp**  
  Allows users to enroll a wallet address and Discord account to establish a secure mapping for the Discord bot to reference.

- 🤖 **Discord Bot**  
  Sends verification requests and allows users to confirm or deny transactions directly through Discord DMs.

- 🧠 **Custom Detector Logic**  
  Evaluates the context and risk of each transaction and determines when to trigger the 2FA flow via Discord.

---
