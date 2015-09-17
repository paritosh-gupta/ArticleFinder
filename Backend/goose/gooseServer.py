import json
import cherrypy
from goose import Goose

class ArticleExtractor(object):
    @cherrypy.expose
    def index(self,url=""):
        if url=="" :
            return "null"
        else:
            print url
            g=Goose()
            article =g.extract(url=url)
            data={'Title':article.title,'Description':article.meta_description,'Content':article.cleaned_text[:],'Image':article.top_image.src}
            return json.dumps(data, sort_keys=True, \
                 indent=4, separators=(',', ': '))

if __name__ == '__main__':
    cherrypy.config.update({'server.socket_port': 8081})
    cherrypy.quickstart(ArticleExtractor())
