mport random
import string
import json
url="http://vrfocus.com/archives/21132/samsung-partner-with-bmw-for-gear-vr-test-drive/"
# import cherrypy
from goose import Goose
g=Goose()
article =g.extract(url=url)
data={'Title':article.title,'Description':article.meta_description,'Content':article.cleaned_text[:],'Image':article.top_image.src}
print json.dumps(data, sort_keys=True, \
indent=4, separators=(',', ': '))
