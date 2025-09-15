# Project PullUp 🚀
**Apply smart, not hard.** Upload your resume, pick target companies (or use a curated default set), and get a **company-wise, ranked list** of roles that match your skills and YOE—no more blind mass-applications.

---

## What it does (MVP)
- 🧠 **Understands your resume:** Extracts skills + **Years of Experience** from dates.
- 🔎 **Finds real roles:** Pulls from Greenhouse/Lever/company careers (FAANG + OpenAI, Airbnb, Spotify, PayPal, Stripe, Intuit, etc. by default).
- 🎯 **Ranks relevance:** Title/skill overlap + seniority vs YOE → **sorted, exportable** results (CSV/PDF).

---

## Tech at a glance
- **Frontend:** React (Next.js)  
- **API:** Go + GraphQL (gqlgen)  
- **Workers:** Go (resume parsing / ranking), Node (job ingestors)  
- **Data:** **DynamoDB-first** (free tier), S3 for file storage  
- **Infra (free-lean):** API Gateway + Lambda, S3, DynamoDB; containerize later (k3s/Helm)  
- **CI/CD:** GitHub Actions → GHCR (later)

---

## High-level flow
1. **Upload resume** → stored in S3 → parse text, compute **YOE**, extract **skills**.  
2. **Ingest jobs** from selected companies → normalize and store in DynamoDB.  
3. **Rank** jobs per resume terms → return **company-grouped** results → **export**.


