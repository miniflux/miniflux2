// Code generated by go generate; DO NOT EDIT.
// 2018-01-16 11:17:12.830111 +0100 CET m=+0.026002079

package locale

var translations = map[string]string{
	"de_DE": `{
    "plural.feed.error_count": [
        "%d Error",
        "%d Errors"
    ],
    "plural.categories.feed_count": [
        "Es gibt %d Abonnement.",
        "Es gibt %d Abonnements."
    ],
    "Username": "Benutzername",
    "Password": "Passwort",
    "Unread": "Ungelesen",
    "History": "Verlauf",
    "Feeds": "Abonnements",
    "Categories": "Kategorien",
    "Settings": "Einstellungen",
    "Logout": "Abmelden",
    "Next": "Nächster Eintrag",
    "Previous": "Vorheriger Eintrag",
    "New Subscription": "Neues Abonnement",
    "Import": "Importieren",
    "Export": "Exportieren",
    "There is no category. You must have at least one category.": "Es ist keine Kategorie vorhanden. Wenigstens eine Kategorie muss angelegt sein.",
    "URL": "URL",
    "Category": "Kategorie",
    "Find a subscription": "Abonnement finden",
    "Loading...": "Lade...",
    "Create a category": "Kategorie anlegen",
    "There is no category.": "Es ist keine Kategorie vorhanden.",
    "Edit": "Bearbeiten",
    "Remove": "Entfernen",
    "No feed.": "Kein Abonnement.",
    "There is no article in this category.": "Es befindet sich kein Artikel in dieser Kategorie.",
    "Original": "Original-Link",
    "Mark this page as read": "Diese Seite als gelesen markieren",
    "not yet": "noch nicht",
    "just now": "gerade",
    "1 minute ago": "vor einer Minute",
    "%d minutes ago": "vor %d Minuten",
    "1 hour ago": "vor einer Stunde",
    "%d hours ago": "vor %d Stunden",
    "yesterday": "gestern",
    "%d days ago": "vor %d Tagen",
    "%d weeks ago": "vor %d Wochen",
    "%d months ago": "vor %d Monaten",
    "%d years ago": "vor %d Jahren",
    "Date": "Datum",
    "IP Address": "IP Adresse",
    "User Agent": "User Agent",
    "Actions": "Aktionen",
    "Current session": "Aktuelle Sitzung",
    "Sessions": "Sitzungen",
    "Users": "Benutzer",
    "Add user": "Benutzer anlegen",
    "Choose a Subscription": "Abonnement auswählen",
    "Subscribe": "Abonnieren",
    "New Category": "Neue Kategorie",
    "Title": "Titel",
    "Save": "Speichern",
    "or": "oder",
    "cancel": "abbrechen",
    "New User": "Neuer Benutzer",
    "Confirmation": "Bestätigung",
    "Administrator": "Administrator",
    "Edit Category: %s": "Kategorie bearbeiten : %s",
    "Update": "Aktualisieren",
    "Edit Feed: %s": "Abonnement bearbeiten : %s",
    "There is no category!": "Es ist keine Kategorie vorhanden!",
    "Edit user: %s": "Benutzer bearbeiten: %s",
    "There is no article for this feed.": "Es existiert kein Artikel für dieses Abonnement.",
    "Add subscription": "Abonnement hinzufügen",
    "You don't have any subscription.": "Es sind keine Abonnements vorhanden",
    "Last check:": "Letzte Aktualisierung:",
    "Refresh": "Aktualisieren",
    "There is no history at the moment.": "Es exisitiert zur Zeit kein Verlauf.",
    "OPML file": "OPML Datei",
    "Sign In": "Anmeldung",
    "Sign in": "Anmelden",
    "Theme": "Thema",
    "Timezone": "Zeitzone",
    "Language": "Sprache",
    "There is no unread article.": "Es existiert kein ungelesener Artikel.",
    "You are the only user.": "Sie sind der einzige Benutzer.",
    "Last Login": "Letzte Anmeldung",
    "Yes": "Ja",
    "No": "Nein",
    "This feed already exists (%s).": "Diese Abonnement existiert bereits (%s).",
    "Unable to fetch feed (statusCode=%d).": "Abonnement konnte nicht abgerufen werden (code=%d).",
    "Unable to open this link: %v": "Dieser Link konnte nicht geöffnet werden: %v",
    "Unable to analyze this page: %v": "Diese Seite konnte nicht analysiert werden: %v",
    "Unable to find any subscription.": "Es wurden keine Abonnements gefunden.",
    "The URL and the category are mandatory.": "Die URL und die Kategorie sind obligatorisch.",
    "All fields are mandatory.": "Alle Felder sind obligatorisch.",
    "Passwords are not the same.": "Passwörter stimmen nicht überein.",
    "You must use at least 6 characters.": "Wenigstens 6 Zeichen müssen genutzt werden.",
    "The username is mandatory.": "Der Benutzername ist obligatorisch.",
    "The username, theme, language and timezone fields are mandatory.": "Die Felder für Benutzername, Thema, Sprache und Zeitzone sind obligatorisch.",
    "The title is mandatory.": "Der Titel ist obligatorisch.",
    "About": "Über",
    "version": "Version",
    "Version:": "Version :",
    "Build Date:": "Datum der Kompilierung:",
    "Author:": "Autor:",
    "Authors": "Autoren",
    "License:": "Lizenz:",
    "Attachments": "Anhänge",
    "Download": "Herunterladen",
    "Invalid username or password.": "Benutzername oder Passwort ungültig.",
    "Never": "Niemals",
    "Unable to execute request: %v": "Diese Anfrage konnte nicht ausgeführt werden: %v",
    "Last Parsing Error": "Letzter Analyse Error",
    "There is a problem with this feed": "Es gibt ein Problem mit diesem Abonnement",
    "Unable to parse OPML file: %v.": "OPML Datei konnte nicht gelesen werden: %v.",
    "Unable to parse RSS feed: %v.": "RSS Feed konnte nicht gelesen werden: %v.",
    "Unable to parse Atom feed: %v.": "Atom Feed konnte nicht gelesen werden: %v.",
    "Unable to parse JSON feed: %v.": "JSON Feed konnte nicht gelesen werden: %v.",
    "Unable to parse RDF feed: %v.": "RDF Feed konnte nicht gelesen werden: %v.",
    "Unable to normalize encoding: %v.": "Zeichenkodierung konnte nicht normalisiert werden: %v.",
    "Unable to create this category.": "Diese Kategorie konnte nicht angelegt werden.",
    "yes": "ja",
    "no": "nein",
    "Are you sure?": "Sind Sie sicher?",
    "Work in progress...": "Laufende Arbeit...",
    "This user already exists.": "Dieser Benutzer existiert bereits.",
    "This category already exists.": "Diese Kategorie existiert bereits.",
    "Unable to update this category.": "Diese Kategorie konnte nicht aktualisiert werden.",
    "Integrations": "Externe Dienste",
    "Bookmarklet": "Bookmarklet",
    "Drag and drop this link to your bookmarks.": "Ziehen Sie diesen Link in Ihre Lesezeichen.",
    "This special link allows you to subscribe to a website directly by using a bookmark in your web browser.": "Dieser spezielle Link ermöglicht es, eine Webseite direkt über ein Lesezeichen im Browser zu abonnieren.",
    "Add to Miniflux": "Mit Miniflux abonnieren",
    "Refresh all feeds in background": "Alle Abonnements im Hintergrund aktualisieren",
    "Sign in with Google": "Mit Google anmelden",
    "Unlink my Google account": "Google Konto abmelden",
    "Link my Google account": "Google Konto assoziieren",
    "Category not found for this user.": "Diese Kategorie existiert nicht für diesen Benutzer.",
    "Invalid theme.": "Dieses Thema ist fehlerhaft.",
    "Entry Sorting": "Sortierung der Einträge",
    "Older entries first": "Ältere Einträge zuerst",
    "Recent entries first": "Neuste Einträge zuerst",
    "Saving...": "Speichern...",
    "Done!": "Erledigt!",
    "Save this article": "Diesen Artikel speichern",
    "Mark bookmark as unread": "Lesezeichen als ungelesen markieren",
    "Pinboard Tags": "Pinboard Tags",
    "Pinboard API Token": "Pinboard API Token",
    "Save articles to Pinboard": "Artikel in Pinboard speichern",
    "Save articles to Instapaper": "Artikel in Instapaper speichern",
    "Instapaper Username": "Instapaper Benutzername",
    "Instapaper Password": "Instapaper Passwort",
    "Activate Fever API": "Fever API aktivieren",
    "Fever Username": "Fever Benutzername",
    "Fever Password": "Fever Passwort",
    "Fetch original content": "Inhalt herunterladen",
    "Scraper Rules": "Extraktionsregeln",
    "Rewrite Rules": "Umschreiberegeln",
    "Preferences saved!": "Einstellungen gespeichert!",
    "Your external account is now linked !": "Ihr externes Konto wurde verlinkt!",
    "Save articles to Wallabag": "Artikel in Wallabag speichern",
    "Wallabag API Endpoint": "Wallabag URL",
    "Wallabag Client ID": "Wallabag Client-ID",
    "Wallabag Client Secret": "Wallabag Client-Secret",
    "Wallabag Username": "Wallabag Benutzername",
    "Wallabag Password": "Wallabag Passwort",
    "Keyboard Shortcut: %s": "Tastenkürzel: %s",
    "Favorites": "Favoriten",
    "Star": "Markieren",
    "Unstar": "Markierung entfernen",
    "Starred": "Markiert",
    "There is no bookmark at the moment.": "Es existiert derzeit kein Lesezeichen.",
    "Last checked:": "Zuletzt geprüft:",
    "ETag header:": "ETag Header:",
    "LastModified header:": "Zuletzt geändert:",
    "None": "Nicht verfügbar",
    "Keyboard Shortcuts": "Tastenkürzel",
    "Sections Navigation": "Navigation zwischen den Sektionen",
    "Go to unread": "Zu den ungelesenen Einträgen gehen",
    "Go to bookmarks": "Zu den Lesezeichen gehen",
    "Go to history": "Zum Verlauf gehen",
    "Go to feeds": "Zu den Abonnements gehen",
    "Go to categories": "Zu den Kategorien gehen",
    "Go to settings": "Zu den Einstellungen gehen",
    "Show keyboard shortcuts": "Tastenkürzel anzeigen",
    "Items Navigation": "Navigation zwischen den Elementen",
    "Go to previous item": "Zum vorherigen Eintrag gehen",
    "Go to next item": "Zum nächsten Eintrag gehen",
    "Go to previous page": "Zur vorherigen Seite gehen",
    "Go to next page": "Zur nächsten Seite gehen",
    "Open selected item": "Gewählten Eintrag öffnen",
    "Open original link": "Original-Link öffnen",
    "Toggle read/unread": "Umschalten zwischen gelesen/ungelesen",
    "Mark current page as read": "Aktuelle Seite als gelesen markieren",
    "Download original content": "Vollständigen Inhalt herunterladen",
    "Toggle bookmark": "Lesezeichen hinzufügen/entfernen",
    "Close modal dialog": "Modalen Dialog schließen",
    "Save article": "Artikel speichern",
    "There is already someone associated with this provider!": "Es ist bereits jemand mit diesem Anbieter assoziiert!",
    "There is already someone else with the same Fever username!": "Es existiert bereits jemand mit diesem Fever Benutzernamen!",
    "Mark all as read": "Alle als gelesen markieren",
    "This feed is empty": "Dieses Abonnement ist leer",
    "Flush history": "Verlauf leeren",
    "Site URL": "Webseite-URL",
    "Feed URL": "Feed-URL"
}
`,
	"en_US": `{
    "plural.feed.error_count": [
        "%d error",
        "%d errors"
    ],
    "plural.categories.feed_count": [
        "There is %d feed.",
        "There are %d feeds."
    ]
}`,
	"fr_FR": `{
    "plural.feed.error_count": [
        "%d erreur",
        "%d erreurs"
    ],
    "plural.categories.feed_count": [
        "Il y a %d abonnement.",
        "Il y a %d abonnements."
    ],
    "Username": "Nom d'utilisateur",
    "Password": "Mot de passe",
    "Unread": "Non lus",
    "History": "Historique",
    "Feeds": "Abonnements",
    "Categories": "Catégories",
    "Settings": "Réglages",
    "Logout": "Se déconnecter",
    "Next": "Suivant",
    "Previous": "Précédent",
    "New Subscription": "Nouvel Abonnment",
    "Import": "Importation",
    "Export": "Exportation",
    "There is no category. You must have at least one category.": "Il n'y a aucune catégorie. Vous devez avoir au moins une catégorie.",
    "URL": "URL",
    "Category": "Catégorie",
    "Find a subscription": "Trouver un abonnement",
    "Loading...": "Chargement...",
    "Create a category": "Créer une catégorie",
    "There is no category.": "Il n'y a aucune catégorie.",
    "Edit": "Modifier",
    "Remove": "Supprimer",
    "No feed.": "Aucun abonnement.",
    "There is no article in this category.": "Il n'y a aucun article dans cette catégorie.",
    "Original": "Original",
    "Mark this page as read": "Marquer cette page comme lu",
    "not yet": "pas encore",
    "just now": "à l'instant",
    "1 minute ago": "il y a une minute",
    "%d minutes ago": "il y a %d minutes",
    "1 hour ago": "il y a une heure",
    "%d hours ago": "il y a %d heures",
    "yesterday": "hier",
    "%d days ago": "il y a %d jours",
    "%d weeks ago": "il y a %d semaines",
    "%d months ago": "il y a %d mois",
    "%d years ago": "il y a %d ans",
    "Date": "Date",
    "IP Address": "Adresse IP",
    "User Agent": "Navigateur Web",
    "Actions": "Actions",
    "Current session": "Session actuelle",
    "Sessions": "Sessions",
    "Users": "Utilisateurs",
    "Add user": "Ajouter un utilisateur",
    "Choose a Subscription": "Choisissez un abonnement",
    "Subscribe": "S'abonner",
    "New Category": "Nouvelle Catégorie",
    "Title": "Titre",
    "Save": "Sauvegarder",
    "or": "ou",
    "cancel": "annuler",
    "New User": "Nouvel Utilisateur",
    "Confirmation": "Confirmation",
    "Administrator": "Administrateur",
    "Edit Category: %s": "Modification de la catégorie : %s",
    "Update": "Mettre à jour",
    "Edit Feed: %s": "Modification de l'abonnement : %s",
    "There is no category!": "Il n'y a aucune catégorie !",
    "Edit user: %s": "Modification de l'utilisateur : %s",
    "There is no article for this feed.": "Il n'y a aucun article pour cet abonnement.",
    "Add subscription": "Ajouter un abonnement",
    "You don't have any subscription.": "Vous n'avez aucun abonnement",
    "Last check:": "Dernière vérification :",
    "Refresh": "Actualiser",
    "There is no history at the moment.": "Il n'y a aucun historique pour le moment.",
    "OPML file": "Fichier OPML",
    "Sign In": "Connexion",
    "Sign in": "Connexion",
    "Theme": "Thème",
    "Timezone": "Fuseau horaire",
    "Language": "Langue",
    "There is no unread article.": "Il n'y a rien de nouveau à lire.",
    "You are the only user.": "Vous êtes le seul utilisateur.",
    "Last Login": "Dernière connexion",
    "Yes": "Oui",
    "No": "Non",
    "This feed already exists (%s).": "Cet abonnement existe déjà (%s).",
    "Unable to fetch feed (statusCode=%d).": "Impossible de récupérer cet abonnement (code=%d).",
    "Unable to open this link: %v": "Impossible d'ouvrir ce lien : %v",
    "Unable to analyze this page: %v": "Impossible d'analyzer cette page : %v",
    "Unable to find any subscription.": "Impossible de trouver un abonnement.",
    "The URL and the category are mandatory.": "L'URL et la catégorie sont obligatoire.",
    "All fields are mandatory.": "Tous les champs sont obligatoire.",
    "Passwords are not the same.": "Les mots de passe ne sont pas les mêmes.",
    "You must use at least 6 characters.": "Vous devez utiliser au moins 6 caractères.",
    "The username is mandatory.": "Le nom d'utilisateur est obligatoire.",
    "The username, theme, language and timezone fields are mandatory.": "Le nom d'utilisateur, le thème, la langue et le fuseau horaire sont obligatoire.",
    "The title is mandatory.": "Le titre est obligatoire.",
    "About": "A propos",
    "version": "Version",
    "Version:": "Version :",
    "Build Date:": "Date de la compilation :",
    "Author:": "Auteur :",
    "Authors": "Auteurs",
    "License:": "Licence :",
    "Attachments": "Pièces jointes",
    "Download": "Télécharger",
    "Invalid username or password.": "Mauvais identifiant ou mot de passe.",
    "Never": "Jamais",
    "Unable to execute request: %v": "Impossible d'exécuter cette requête: %v",
    "Last Parsing Error": "Dernière erreur d'analyse",
    "There is a problem with this feed": "Il y a un problème avec cet abonnement",
    "Unable to parse OPML file: %v.": "Impossible de lire ce fichier OPML : %v.",
    "Unable to parse RSS feed: %v.": "Impossible de lire ce flux RSS: %v.",
    "Unable to parse Atom feed: %v.": "Impossible de lire ce flux Atom: %v.",
    "Unable to parse JSON feed: %v.": "Impossible de lire ce flux JSON: %v.",
    "Unable to parse RDF feed: %v.": "Impossible de lire ce flux RDF: %v.",
    "Unable to normalize encoding: %v.": "Impossible de normaliser l'encodage : %v.",
    "Unable to create this category.": "Impossible de créer cette catégorie.",
    "yes": "oui",
    "no": "non",
    "Are you sure?": "Êtes-vous sûr ?",
    "Work in progress...": "Travail en cours...",
    "This user already exists.": "Cet utilisateur existe déjà.",
    "This category already exists.": "Cette catégorie existe déjà.",
    "Unable to update this category.": "Impossible de mettre à jour cette catégorie.",
    "Integrations": "Intégrations",
    "Bookmarklet": "Bookmarklet",
    "Drag and drop this link to your bookmarks.": "Glisser-déposer ce lien dans vos favoris.",
    "This special link allows you to subscribe to a website directly by using a bookmark in your web browser.": "Ce lien spécial vous permet de vous abonner à un site web directement en utilisant un marque page dans votre navigateur web.",
    "Add to Miniflux": "Ajouter à Miniflux",
    "Refresh all feeds in background": "Actualiser les abonnements en arrière-plan",
    "Sign in with Google": "Se connecter avec Google",
    "Unlink my Google account": "Dissocier mon compte Google",
    "Link my Google account": "Associer mon compte Google",
    "Category not found for this user.": "Cette catégorie n'existe pas pour cet utilisateur.",
    "Invalid theme.": "Le thème est invalide.",
    "Entry Sorting": "Ordre des éléments",
    "Older entries first": "Ancien éléments en premier",
    "Recent entries first": "Éléments récents en premier",
    "Saving...": "Enregistrement...",
    "Done!": "Terminé !",
    "Save this article": "Sauvegarder cet article",
    "Mark bookmark as unread": "Marquer le lien comme non lu",
    "Pinboard Tags": "Libellés de Pinboard",
    "Pinboard API Token": "Jeton de sécurité de l'API de Pinboard",
    "Save articles to Pinboard": "Sauvegarder les articles vers Pinboard",
    "Save articles to Instapaper": "Sauvegarder les articles vers Instapaper",
    "Instapaper Username": "Nom d'utilisateur Instapaper",
    "Instapaper Password": "Mot de passe Instapaper",
    "Activate Fever API": "Activer l'API de Fever",
    "Fever Username": "Nom d'utilisateur pour l'API de Fever",
    "Fever Password": "Mot de passe pour l'API de Fever",
    "Fetch original content": "Récupérer le contenu original",
    "Scraper Rules": "Règles pour récupérer le contenu original",
    "Rewrite Rules": "Règles de réécriture",
    "Preferences saved!": "Préférences sauvegardées !",
    "Your external account is now linked !": "Votre compte externe est maintenant associé !",
    "Save articles to Wallabag": "Sauvegarder les articles vers Wallabag",
    "Wallabag API Endpoint": "URL de l'API de Wallabag",
    "Wallabag Client ID": "Identifiant du client Wallabag",
    "Wallabag Client Secret": "Clé secrète du client Wallabag",
    "Wallabag Username": "Nom d'utilisateur de Wallabag",
    "Wallabag Password": "Mot de passe de Wallabag",
    "Keyboard Shortcut: %s": "Raccourci clavier : %s",
    "Favorites": "Favoris",
    "Star": "Favoris",
    "Unstar": "Enlever favoris",
    "Starred": "Favoris",
    "There is no bookmark at the moment.": "Il n'y a aucun favoris pour le moment.",
    "Last checked:": "Dernière vérification :",
    "ETag header:": "En-tête ETag :",
    "LastModified header:": "En-tête LastModified :",
    "None": "Aucun/Aucune",
    "Keyboard Shortcuts": "Raccourcis clavier",
    "Sections Navigation": "Naviguation entre les sections",
    "Go to unread": "Aller aux éléments non lus",
    "Go to bookmarks": "Aller aux favoris",
    "Go to history": "Voir l'historique",
    "Go to feeds": "Voir les abonnements",
    "Go to categories": "Voir les catégories",
    "Go to settings": "Voir les réglages",
    "Show keyboard shortcuts": "Voir les raccourcis clavier",
    "Items Navigation": "Naviguation entre les éléments",
    "Go to previous item": "Élément précédent",
    "Go to next item": "Élément suivant",
    "Go to previous page": "Page précédente",
    "Go to next page": "Page suivante",
    "Open selected item": "Ouvrir élément sélectionné",
    "Open original link": "Ouvrir lien original",
    "Toggle read/unread": "Basculer entre lu/non lu",
    "Mark current page as read": "Marquer la page actuelle comme lu",
    "Download original content": "Télécharger le contenu original",
    "Toggle bookmark": "Ajouter/Enlever favoris",
    "Close modal dialog": "Fermer la boite de dialogue",
    "Save article": "Sauvegarder l'article",
    "There is already someone associated with this provider!": "Il y a déjà quelqu'un d'associé avec ce provider !",
    "There is already someone else with the same Fever username!": "Il y a déjà quelqu'un d'autre avec le même nom d'utilisateur Fever !",
    "Mark all as read": "Tout marquer comme lu",
    "This feed is empty": "Cet abonnement est vide"
}
`,
}

var translationsChecksums = map[string]string{
	"de_DE": "93b493c13389bc6d26110f9200ba287593957ae1f97f147ee889024b35962cbc",
	"en_US": "6fe95384260941e8a5a3c695a655a932e0a8a6a572c1e45cb2b1ae8baa01b897",
	"fr_FR": "888baaeba8cee020f0cf77db7bad9960e9b640277d042029cafb41e72e13566e",
}
