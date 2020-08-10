# -*- coding: utf-8 -*-
import re
import sys

from workflow import web, Workflow3


class StackOverflow(object):
    BASE_URL = "https://stackoverflow.com/search"
    DOAMIN = "https://stackoverflow.com"
    ICON = './icon.png'
    HEADERS = {
        'User-Agent':
            "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36",
        'referer': "https://stackoverflow.com/",
        'sec-fetch-site': "same-origin",
        'Accept': "*/*",
        'Cache-Control': "no-cache",
        'Host': "stackoverflow.com",
        'Accept-Encoding': "gzip, deflate",
    }
    CHARACTER_ENTITIES = {
        '&ensp;': ' ',
        '&emsp;': ' ',
        '&nbsp;': ' ',
        '&lt;': '<',
        '&gt;': '>',
        '&amp;': '&',
        '&#39;': '\'',
        '&quot;': '"',
        "&ldquo;": '“',
        "&rdquo;": '”',
        '&copy;': '©',
        '&reg;': '®',
        '™': '™',
        '&times;': '×',
        '&divide;': '÷'
    }

    def __init__(self, wf):
        self.session = web
        self.params = {"q": wf.args[0], 'tab': 'votes'}
        self.res = []
        self.wf = wf

    def _query_keyword(self):
        res = self.session.get(self.BASE_URL,
                               params=self.params,
                               headers=self.HEADERS)
        return res.content

    def _filter_result(self):
        html_content = self.wf.cached_data(self.wf.args[0], self._query_keyword, max_age=1800)
        vote_pattern = re.compile(r'.*class="vote-count-post "><strong>(\d+)')
        url_pattern = re.compile(r'<a href="(.*?)\?r=SearchResults#')
        title_pattern = re.compile(
            r'class="question-hyperlink">.*?: (.*?)\s*</a>', re.M | re.S)
        titles = title_pattern.findall(html_content, pos=30000)
        urls = url_pattern.findall(html_content, pos=30000)
        votes = vote_pattern.findall(html_content, pos=30000)
        for i in zip(titles, urls, votes):
            self.res.append({'title': i[0], 'url': i[1], 'vote': i[2]})

    def _format_result(self):
        for item in self.res:
            for char in self.CHARACTER_ENTITIES:
                if char in item['title']:
                    item['title'] = item["title"].replace(
                        char, self.CHARACTER_ENTITIES[char])

            self.wf.add_item(title=item['title'],
                             subtitle='Votes: %s' % item['vote'],
                             icon=self.ICON,
                             arg=self.DOAMIN + item['url'],
                             valid=True)
        self.wf.send_feedback()


def main(wf):
    stack = StackOverflow(wf)
    stack._filter_result()
    stack._format_result()


if __name__ == '__main__':
    wf = Workflow3()
    sys.exit(wf.run(main))
