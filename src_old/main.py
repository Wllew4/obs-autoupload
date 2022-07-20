import time
import os
import shutil
from .VOD import VOD
from .yt_upload import upload
from .yt_playlist import add_to_playlist

from .secrets import CONFIG

from PyQt5.QtWidgets import QApplication, QWidget, QLabel, QPushButton, QVBoxLayout
import sys
import threading

def main():

	class Window(QWidget):
		def cleanup(self):
			print('Starting cleanup')
			# Cleanup
			if CONFIG['files']['POST_UPLOAD_STEP'] == "delete":
				os.remove(self.vod.path)
			if CONFIG['files']['POST_UPLOAD_STEP'] == "move":
				shutil.move(
					self.vod.path,
					CONFIG['files']['VOD_ARCHIVE_DIR'] + '/' + os.path.basename(self.vod.path))
				open(
					CONFIG['files']['VOD_ARCHIVE_DIR']
					+ '\\'
					+ os.path.basename(self.vod.path)
					+ '.meta', 'w', encoding='utf-8').write('%s\n%s\n%s\n%s\n' % (
						self.vod.title, self.vod.date, self.vod.yt_url, self.vod.ttv_url
					))

			# Congrats
			self.label.setText('Cleanup complete!')
			self.btn_proceed.setParent(None)
			self.btn_cancel.setText('Close')

		# Upload VOD
		def upload(self):
			print('Starting upload')
			self.btn_proceed.setText('Confirm Cleanup')
			self.btn_proceed.clicked.connect(self.cleanup)
			self.btn_proceed.clicked.disconnect(self.upload)
			LOG = 'Uploading...'
			self.label.setText(LOG)

			# Upload VOD
			self.vod.yt_url = upload(self.vod.path, self.vod.title,
				CONFIG['upload']['DESCRIPTION'].format(
					TITLE=self.vod.title,
					DATE=self.vod.date,
					URL=self.vod.ttv_url,
					ID=self.vod.stream_id,
					PATH=self.vod.path
				),
				CONFIG['upload']['VISIBILITY'],
				CONFIG['upload']['CATEGORY_ID'],
				CONFIG['upload']['TAGS'])
			if type(self.vod.yt_url) is tuple:
				LOG += '\nFailed to upload\n%s' % self.vod.yt_url[1]
				self.label.setText(LOG)
				self.btn_proceed.setParent(None)
			else:
				LOG += '\nUpload Complete!'
				self.label.setText(LOG)

				# Add video to playlist
				if CONFIG['upload']['PLAYLIST_ID']:
					LOG += '\nAdding to playlist...'
					self.label.setText(LOG)
					time.sleep(2)
					add_to_playlist(CONFIG['upload']['PLAYLIST_ID'], self.vod.yt_url)
					LOG += '\nAdded to playlist!'
					self.label.setText(LOG)
				
				# Confirm cleanup settings
				LOG += '\nConfirm cleanup settings:\nMODE:\t%s' % CONFIG['files']['POST_UPLOAD_STEP']
				if CONFIG['files']['POST_UPLOAD_STEP'] == "move":
					LOG += "\nMove to:\t%s" % CONFIG['files']['VOD_ARCHIVE_DIR'] + '/' + os.path.basename(self.vod.path)
				self.label.setText(LOG)

		# Fetch VOD info
		def fetch_info(self):
			self.label.setText('Fetching VOD info...')

			self.vod: VOD = VOD()
			print(self.vod.title, self.vod.date, self.vod.ttv_url, self.vod.stream_id, self.vod.path, sep='\n')

			if self.vod.path == '':
				self.label.setText('VOD file not found')
				self.btn_proceed.setParent(None)
				self.btn_cancel.setText('Close')
			else:
				self.label.setText('Confirm VOD info:\nTITLE:\t%s\nDATE:\t%s\nTTV URL:\t%s\nPATH:\t%s'
					% (self.vod.title, self.vod.date, self.vod.ttv_url, self.vod.path))
				self.btn_proceed.setText('Confirm Upload')
				self.btn_proceed.clicked.connect(self.upload)

		def __init__(self):
			# Initialize window
			super().__init__()
			self.setWindowTitle("OBS Auto-Upload")
			self.setGeometry(0, 0, 400, 200)
			self.vbox = QVBoxLayout(self)

			# Label
			self.label = QLabel('')
			self.label.setWordWrap(True)
			self.vbox.addWidget(self.label)
			# Confirm button
			self.btn_proceed = QPushButton("...")
			self.vbox.addWidget(self.btn_proceed)
			# Cancel button
			self.btn_cancel = QPushButton("Cancel")
			self.btn_cancel.clicked.connect(self.close)
			self.vbox.addWidget(self.btn_cancel)

			# Apply layout
			self.setLayout(self.vbox)
			self.show()
			# Start workflow
			threading.Thread(target=self.fetch_info).start()
			

	App = QApplication(sys.argv)
	window = Window()
	App.exec()
