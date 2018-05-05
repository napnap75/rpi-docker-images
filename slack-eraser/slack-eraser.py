import argparse
import configparser
import time
import re

from slacker import Slacker

# Parse the arguments
parser = argparse.ArgumentParser(description='Utilityy to delete Slack message.')
parser.add_argument('file', help='The configuration file')
parser.add_argument('--verbose', '-v', action='count', help='Be verbose')
parser.add_argument('--dry-run', '-n', action='count', help='Do nothing (for test purpose)')
args = parser.parse_args()
if args.verbose > 0 :
	print("Using args:", args)

# Parse the configuration file
if args.verbose > 0 :
	print("Parsing configuration file:", args.file)
config = configparser.ConfigParser()
config.read(args.file)

# Open the connection to Slack
if args.verbose > 0 :
	print("Opening Slack connection")
slack = Slacker(config['DEFAULT']['SlackTocken'])

# Iterate over all sections in the config file
for sectionName in config.sections():
	section = config[sectionName]
	eraserType = section.get('Type')

	searchRequest = section.get('Request', "")
	searchFrom = section.get('From')
	if searchFrom != None and " " not in searchFrom:
		searchRequest += " from:" + searchFrom
	searchIn = section.get('In')
	if searchIn != None:
		searchRequest += " in:" + searchIn
	searchSort = section.get('Sort', "timestamp")
	searchOrder = section.get('Order', "desc")
	searchCount = section.get('Count', 100)
	searchRE = section.get('RegExp')
	regexp = None
	if searchRE != None:
		regexp = re.compile(searchRE)

	if eraserType == 'duplicates':
		if args.verbose > 0 :
			print("Deleting duplicates messages with request: \"", searchRequest, "\", sort:", searchSort, ", order:", searchOrder, "and count:", searchCount)
			if searchRE != None:
				print("And using regular expression:", searchRE)

		searchResponse = slack.search.messages(searchRequest, searchSort, searchOrder, None, searchCount)
		messages = searchResponse.body['messages']['matches']
		seenMessages = []
		for message in messages:
			text = message['text']
			if text == "":
				try:
					text = message['attachments'][0]['fallback']
				except KeyError as ex:
					try:
						text = message['attachments'][0]['text']
					except KeyError as ex:
						if args.verbose > 1:
							print("Message skipped because empty")
						continue
			if searchFrom != None and searchFrom != message['username']:
				if args.verbose > 1:
					print(text, "--> Skipped (From)")
				continue
			if searchIn != None and searchIn != message['channel']['name']:
				if args.verbose > 1:
					print(text, "--> Skipped (In)")
				continue

			key = text
			if regexp != None:
				match = regexp.match(text)
				if match:
					try:
						key = match.group(1)
					except IndexError as ex:
						key = match.group(0)
					if args.verbose > 1:
						print(text, "using key:", key)
				else:
					if args.verbose > 1:
						print(key, "--> Skipped (RegExp)")
					continue

			if key in seenMessages:
				if args.verbose > 1:
					print(key, "--> Already seen")
				if args.dry_run == None:
					if args.verbose > 0:
						print("Deleting message timestamp", message['ts'], "in channel", message['channel']['name'])
					slack.chat.delete(message['channel']['id'], message['ts'])
					time.sleep(60/50)
				else:
					if args.verbose > 0:
						print("Should have deleted message timestamp", message['ts'], "in channel", message['channel']['name'])
			else:
				if args.verbose > 1:
					print(key, "--> Not seen")
				seenMessages.append(key)
	elif eraserType == 'alerts':
		closedStatus = section.get('ClosedStatus', "OK")
		openStatus = section.get('OpenStatus', "CRITICAL")
		if args.verbose > 0 :
			print("Deleting alerts messages with request: \"", searchRequest, "\", sort:", searchSort, ", order:", searchOrder, "and count:", searchCount)
			print("And using regular expression:", searchRE, "with opening status:", openStatus, "and closing status:", closedStatus)

		searchResponse = slack.search.messages(searchRequest, searchSort, searchOrder, None, searchCount)
		messages = searchResponse.body['messages']['matches']
		closedAlerts = []
		for message in messages:
			text = message['text']
			if text == "":
				try:
					text = message['attachments'][0]['fallback']
				except KeyError as ex:
					try:
						text = message['attachments'][0]['text']
					except KeyError as ex:
						if args.verbose > 1:
							print("Message skipped because empty")
						continue
			if searchFrom != None and searchFrom != message['username']:
				if args.verbose > 1:
					print(text, "--> Skipped (From)")
				continue
			if searchIn != None and searchIn != message['channel']['name']:
				if args.verbose > 1:
					print(text, "--> Skipped (In)")
				continue

			status = None
			match = regexp.match(text)
			if match:
				if match.group(1) == closedStatus:
					closedAlerts.append(match.group(2))
					if args.verbose > 1:
						print(text, "--> Alert '", match.group(2), "' closed")
				elif match.group(1) == openStatus:
					if match.group(2) in closedAlerts:
						if args.verbose > 1:
							print(text, "--> Alert '", match.group(2), "' already closed")
					else:
						if args.verbose > 1:
							print(text, "--> Alert '", match.group(2), "' not closed")
						continue
				elif match.group(2) == closedStatus:
					closedAlerts.append(match.group(1))
					if args.verbose > 1:
						print(text, "--> Alert '", match.group(1), "' closed")
				elif match.group(2) == openStatus:
					if match.group(1) in closedAlerts:
						if args.verbose > 1:
							print(text, "--> Alert '", match.group(1), "' already closed")
					else:
						if args.verbose > 1:
							print(text, "--> Alert '", match.group(1), "' not closed")
						continue
				else:
					if args.verbose > 1:
						print(text, "--> Skipped (Status not found)")
					continue
			else:
				if args.verbose > 1:
					print(text, "--> Skipped (Does not match regular expression)")
				continue

			if args.dry_run == None:
				if args.verbose > 0:
					print("Deleting message timestamp", message['ts'], "in channel", message['channel']['name'])
				slack.chat.delete(message['channel']['id'], message['ts'])
				time.sleep(60/50)
			else:
				if args.verbose > 0:
					print("Should have deleted message timestamp", message['ts'], "in channel", message['channel']['name'])
	else:
		print("Unknown type ", eraserType, " in section", sectionName)

