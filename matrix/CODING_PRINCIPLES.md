# Coding Principles — Quick Reference

## Simplicity

- **KISS (Keep It Simple, Stupid)** — sabse simple solution jo kaam kare, wahi likho, clever/complex mat banao.
- **YAGNI (You Aren't Gonna Need It)** — future ke liye "shayad kaam aayega" soch ke feature mat banao, jab zarurat ho tab banao.
- **DRY (Don't Repeat Yourself)** — same code baar-baar copy mat karo, ek jagah likho aur reuse karo.
- **WET (Write Everything Twice)** — DRY ka anti-pattern; ek-do baar repeat chalta hai, teesri baar abstraction banao (**Rule of Three**).
- **Occam's Razor** — do solutions kaam karte hain to jo simpler hai wahi choose karo.

## Control Flow & Errors

- **Fail-fast** — problem milte hi turant ruk jao (`return`/`break`), aage waste kaam mat karo.
- **Guard Clause** — function/loop ke shuru me bad-case check karke turant exit karo, nested `if` mat banao.
- **Defensive Programming** — bahar se aane wala input (user, file, API) kabhi trust mat karo, use karne se pehle validate karo.
- **Graceful Degradation** — poora system crash hone ke bajaye, jitna ho sake utna kaam karte raho jab kuch fail ho.
- **Idempotency** — ek hi operation ko baar-baar chalane se result same rahe (jaise "set value = 5" — 1 baar chalao ya 10 baar, result same).

## Structure & Design (SOLID — OOP ke liye)

- **Single Responsibility** — har function/class sirf ek kaam kare.
- **Open/Closed** — code extend karna aasan ho (naya feature add), lekin purana code modify na karna pade.
- **Liskov Substitution** — child class ko parent ki jagah use karo to kuch tootna nahi chahiye.
- **Interface Segregation** — bada, sab-kuch-wala interface banane ke bajaye, chhote specific interfaces banao.
- **Dependency Inversion** — concrete implementation pe nahi, abstraction (interface) pe depend karo.

## Code Organization

- **Separation of Concerns** — alag-alag kaam (data fetch, logic, display) alag jagah rakho, sab ek jagah mat mix karo.
- **High Cohesion, Low Coupling** — ek module ke andar related cheezein saath rakho (cohesion), lekin modules ek-dusre pe kam se kam depend karein (coupling).
- **Encapsulation** — internal details chhupao, sirf zaruri cheez bahar expose karo.
- **Law of Demeter** — apne "direct dost" se hi baat karo, kisi cheez ke andar ke andar mat ghuso (`a.b.c.d` jaisi deep chains avoid karo).
- **Composition over Inheritance** — "is-a" relationship (inheritance) ke bajaye "has-a" (composition) prefer karo, zyada flexible hota hai.

## Readability & Habits

- **Principle of Least Astonishment** — code aisa likho jo dusre developer ki expectation ke against surprise na de.
- **Self-documenting Code** — variable/function names itne clear ho ki comments ki zarurat kam pade.
- **Boy Scout Rule** — jo code touch karo, usse thoda clean karke chhodo (jitna mila tha usse behtar).
- **Premature Optimization is the root of all evil** (Donald Knuth) — pehle correct aur simple code likho, speed ka optimization tab karo jab zarurat prove ho.
- **Convention over Configuration** — language/framework ke standard conventions follow karo, har cheez manually configure mat karo.

## Beginner Priority

Sabse pehle inpe focus karo — sabse jaldi impact dikhayenge:
1. **DRY** — repeated code dhundo aur function bana lo
2. **Fail-fast / Guard Clause** — error handling clean karo
3. **Single Responsibility** — chhote, focused functions likho

Baaki sab dheere-dheere apne aap aayenge jaise-jaise zyada code likhoge aur purana code refactor karoge.
