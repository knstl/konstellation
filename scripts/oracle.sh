# set exchange rate
knstld tx oracle set-exchange-rate kethkusd 2709000000 keth,kusd --from hawking --chain-id darchub

# delete exchange rate
knstld tx oracle delete-exchange-rate kbtckusd --from hawking --chain-id darchub

# Update admin addresses set
knstld tx oracle set-admin-addr --add darc1rzdt9wrzwv3x7vv6f7xpyaqqgf3lt6phptqtsx --from hawking --chain-id darchub

# Query all pairs
knstld q oracle all-exchange-rates

# Query one pair
knstld q oracle exchange-rate kbtckusd