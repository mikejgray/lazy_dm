from sys import argv

from PIL import Image, ImageFont, ImageDraw

text = argv[1].upper()

image = Image.open("./public/panel.jpg")
font = ImageFont.truetype('/usr/share/fonts/X11/Type1/c0632bt_.pfb', 80)

image_editable = ImageDraw.Draw(image)
w, h = image.size
fw, fh = font.getsize(text)
cw = (w - fw) / 2
ch = (h - fh) / 2
image_editable.multiline_text((cw, ch), text, fill="white", font=font)
image.save(f"{text.lower()}.jpg")
